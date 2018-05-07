import static java.lang.Integer.parseInt;

class LineCommand implements Command {
    private int x1, y1, x2, y2;
    public LineCommand(String cmd){
        String[] cmdParts = cmd.split("\\s");
        x1 = parseInt(cmdParts[1]);
        y1 = parseInt(cmdParts[2]);
        x2 = parseInt(cmdParts[3]);
        y2 = parseInt(cmdParts[4]);
    }

    public String paint(Coordinate looking) {
        if (x2 >= looking.x && x1 <= looking.x && y2 >= looking.y && y1 <= looking.y) {
            return "X";
        }
        return " ";
    }
}
