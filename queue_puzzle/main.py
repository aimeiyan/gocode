__author__ = 'feng'


class QueueProblem(object):
    def __init__(self, n, pic_path="/tmp/", draw=False):
        self.n = n
        self.result = [-1] * n
        self.pic_index = 0
        self.draw = draw
        self.images = []
        if draw:
            from os import path, mkdir

            self.pic_path = path.join(pic_path, "%squeue" % n)
            if not path.exists(self.pic_path):
                mkdir(self.pic_path)

    def draw_queue(self, n_rows, is_traceback=False):
        from PIL import Image, ImageDraw

        one = 45
        img = Image.new('RGB', (one * self.n + one, one * self.n + one), "#ffffff")
        draw = ImageDraw.Draw(img)

        dark_color = "#d18b47"
        light_color = "#fbc99f"

        colors = [dark_color, light_color]

        for i in range(self.n):
            step = i
            for j in range(self.n):
                draw.rectangle((one * i, one * j, one * i + one, one * j + one), fill=colors[step % 2])
                step += 1

        xs = self.result[:n_rows + 1]
        for i, pos in enumerate(xs):
            radius = one * 0.8
            half = one * 0.1
            y, x = i * one + half, + pos * one + half

            if i == n_rows and is_traceback:
                draw.ellipse((x, y, x + radius, y + radius), fill="#ff0000")
            else:
                draw.ellipse((x, y, x + radius, y + radius), fill="#eee")

        self.pic_index += 1
        self.images.append(img)
        img.save("%s/%s.png" % (self.pic_path, self.pic_index), "png")

    def find_next_legal(self, n_row, trace_back=False):
        result = self.result
        if trace_back:
            res = self.result[n_row] + 1
        else:
            res = 0
        while res < self.n:
            ok = True
            for row in range(n_row):
                if result[row] == res:
                    ok = False
                    break
                if n_row - row == abs(res - result[row]):
                    ok = False
                    break
            if ok:
                return res
            res += 1
        return None

    def solve(self):
        step = 0
        while step < self.n:
            pos = self.find_next_legal(step)

            if pos is not None:
                self.result[step] = pos
                if self.draw: self.draw_queue(step)
                if step == self.n - 1:
                    if self.draw:
                        self.draw = False # only draw the picture for one solution
                        from images2gif import writeGif

                        writeGif("%s/all.gif" % self.pic_path, self.images, duration=0.2, repeat=False)
                    yield self.result

            # start traceback
            if pos is None or step == self.n - 1:
                while 1:
                    step -= 1
                    if step >= 0:
                        pos = self.find_next_legal(step, True)
                        if pos:
                            self.result[step] = pos
                            if self.draw: self.draw_queue(step, is_traceback=True)
                            break
                    else:
                        # all the possibility have been examined
                        return
            step += 1


if __name__ == '__main__':
    queue = QueueProblem(8, draw=True)
    count = 0
    for result in queue.solve():
        count += 1
        print result
        break

    print count

