import urwid
import random

GRID_WIDTH = 4
GRID_HEIGHT = 4

class Board:
    def __init__(self):
        self.grid = [[0] * GRID_WIDTH for _ in range(GRID_HEIGHT)]
        self.score = 0
        self.over = False
        self.add_random_tile()
        self.add_random_tile()

    def add_random_tile(self):
        empty_cells = [(i, j) for i in range(GRID_HEIGHT) for j in range(GRID_WIDTH) if self.grid[i][j] == 0]
        if not empty_cells:
            self.over = True
            return
        i, j = random.choice(empty_cells)
        self.grid[i][j] = 2 if random.random() < 0.9 else 4

    def move_up(self):
        moved = False
        for j in range(GRID_WIDTH):
            for i in range(1, GRID_HEIGHT):
                if self.grid[i][j] != 0:
                    k = i
                    while k > 0 and self.grid[k-1][j] == 0:
                        self.grid[k-1][j] = self.grid[k][j]
                        self.grid[k][j] = 0
                        k -= 1
                        moved = True
                    if k > 0 and self.grid[k-1][j] == self.grid[k][j]:
                        self.grid[k-1][j] *= 2
                        self.score += self.grid[k-1][j]
                        self.grid[k][j] = 0
                        moved = True
        return moved

    def move_down(self):
        moved = False
        for j in range(GRID_WIDTH):
            for i in range(GRID_HEIGHT-2, -1, -1):
                if self.grid[i][j] != 0:
                    k = i
                    while k < GRID_HEIGHT-1 and self.grid[k+1][j] == 0:
                        self.grid[k+1][j] = self.grid[k][j]
                        self.grid[k][j] = 0
                        k += 1
                        moved = True
                    if k < GRID_HEIGHT-1 and self.grid[k+1][j] == self.grid[k][j]:
                        self.grid[k+1][j] *= 2
                        self.score += self.grid[k+1][j]
                        self.grid[k][j] = 0
                        moved = True
        return moved

    def move_left(self):
        moved = False
        for i in range(GRID_HEIGHT):
            for j in range(1, GRID_WIDTH):
                if self.grid[i][j] != 0:
                    k = j
                    while k > 0 and self.grid[i][k-1] == 0:
                        self.grid[i][k-1] = self.grid[i][k]
                        self.grid[i][k] = 0
                        k -= 1
                        moved = True
                    if k > 0 and self.grid[i][k-1] == self.grid[i][k]:
                        self.grid[i][k-1] *= 2
                        self.score += self.grid[i][k-1]
                        self.grid[i][k] = 0
                        moved = True
        return moved

    def move_right(self):
        moved = False
        for i in range(GRID_HEIGHT):
            for j in range(GRID_WIDTH-2, -1, -1):
                if self.grid[i][j] != 0:
                    k = j
                    while k < GRID_WIDTH-1 and self.grid[i][k+1] == 0:
                        self.grid[i][k+1] = self.grid[i][k]
                        self.grid[i][k] = 0
                        k += 1
                        moved = True
                    if k < GRID_WIDTH-1 and self.grid[i][k+1] == self.grid[i][k]:
                        self.grid[i][k+1] *= 2
                        self.score += self.grid[i][k+1]
                        self.grid[i][k] = 0
                        moved = True
        return moved

# Updated palette with valid urwid color names
palette = [
    ('default', 'white', 'black'),
    ('score', 'yellow', 'black'),
    ('gameover', 'dark red', 'black'),
    ('instructions', 'dark green', 'black'),
    ('tile_2', 'dark cyan', 'black'),
    ('tile_4', 'dark blue', 'black'),
    ('tile_8', 'dark green', 'black'),
    ('tile_16', 'yellow', 'black'),
    ('tile_32', 'light red', 'black'),
    ('tile_64', 'dark red', 'black'),
    ('tile_128', 'dark magenta', 'black'),
    ('tile_empty', 'dark gray', 'black'),
]

def get_tile_style(val):
    if val == 0:
        return 'tile_empty'
    elif val == 2:
        return 'tile_2'
    elif val == 4:
        return 'tile_4'
    elif val == 8:
        return 'tile_8'
    elif val == 16:
        return 'tile_16'
    elif val == 32:
        return 'tile_32'
    elif val == 64:
        return 'tile_64'
    else:
        return 'tile_128'

def main():
    board = Board()
    grid_text = urwid.Text("")
    
    def update_display():
        display = [(('score', f"Score: {board.score}"), '\n\n')]
        for i in range(GRID_HEIGHT):
            row = []
            for j in range(GRID_WIDTH):
                val = board.grid[i][j]
                style = get_tile_style(val)
                row.append((style, f"{val:4} " if val != 0 else "    - "))
            display.append(row)
            display.append('\n')
        if board.over:
            display.append(('gameover', '\nGame Over!\n'))
        display.append(('instructions', 'Use arrow keys to move, q to quit'))
        grid_text.set_text(display)

    update_display()
    main_widget = urwid.Filler(grid_text, 'top')
    loop = urwid.MainLoop(main_widget, palette, unhandled_input=lambda key: handle_input(key, board, update_display))
    loop.run()

def handle_input(key, board, update_display):
    if isinstance(key, str):
        if key == 'q':
            raise urwid.ExitMainLoop()
        elif not board.over:
            moved = False
            if key == 'up':
                moved = board.move_up()
            elif key == 'down':
                moved = board.move_down()
            elif key == 'left':
                moved = board.move_left()
            elif key == 'right':
                moved = board.move_right()
            if moved:
                board.add_random_tile()
                update_display()

if __name__ == "__main__":
    main()
