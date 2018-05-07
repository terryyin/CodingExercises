import java.util.ArrayList;
import java.util.List;

import static java.lang.Integer.parseInt;

public class TextPainter {

    public Canvas canvas = new Canvas(10, 10);
    private List<Command> commands = new ArrayList<Command>();

    public TextPainter(int width, int height) {
        this.canvas = new Canvas(width, height);
    }

    public TextPainter command(String cmd) {
        if (cmd.startsWith("R")) {
            String[] lists = cmd.split("\\s");
            String x1 = lists[1];
            String y1 = lists[2];
            String x2 = lists[3];
            String y2 = lists[4];
            this.addCommand(String.join(" ", "L", x1, y1, x2, y1));
            this.addCommand(String.join(" ", "L", x1, y1, x1, y2));
            this.addCommand(String.join(" ", "L", x2, y1, x2, y2));
            this.addCommand(String.join(" ", "L", x1, y2, x2, y2));
        }
        else {
            addCommand(cmd);
        }
        return this;
    }

    private void addCommand(String cmd) {
        if(cmd.startsWith("L")) {
            this.commands.add(new LineCommand(cmd));
            return;
        }
        this.commands.add(new BucketCommand(cmd, dup()));
    }

    private TextPainter dup() {
        TextPainter previousPainter = new TextPainter(canvas.width, canvas.height);
        previousPainter.commands = new ArrayList<>(commands);
        return previousPainter;
    }

    public String pixel(int x, int y) {
        return pixel(new Coordinate(x, y));
    }

    public String pixel(Coordinate cood) {
        return pixelAfterCmds(cood, commands);
    }

    private String pixelAfterCmds(Coordinate p, List<Command> commands) {
        for (Command command : commands) {
            String pix = command.paint(p);
            if (pix != " ")
                return pix;
        }
        return " ";
    }

    public void render() {
        for (int i = 0; i <= canvas.width; i++) {
            for (int j = 0; j <= canvas.height; j++)
                System.out.print(pixel(i, j));
            System.out.print("\n");
        }
    }

    public static void main(String[] args) {
        TextPainter tp = new TextPainter(20, 20);
        tp.command("L 1 2 6 2");
        tp.command("L 6 3 6 20");
        tp.command("R 14 1 18 3");
        tp.command("B 3 3 o");
        tp.render();

    }

}
