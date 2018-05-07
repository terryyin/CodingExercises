import java.util.HashSet;
import java.util.Set;
import java.util.stream.Collectors;

import static java.lang.Integer.parseInt;

class BucketCommand implements Command {
    private final String cmd;
    private final TextPainter previousPainter;
    private Coordinate origin;
    private String color;

    public BucketCommand(String cmd, TextPainter previousPainter){
        this.cmd = cmd;
        this.previousPainter = previousPainter;
        String[] cmdParts = cmd.split("\\s");
        this.origin = new Coordinate(parseInt(cmdParts[1]), parseInt(cmdParts[2]));
        color = cmdParts[3];
    }

    public String paint(Coordinate looking) {
        Set<Coordinate> filled = new HashSet<Coordinate>();
        filled.add(origin);
        return pixelAfterBucketing(looking, filled);
    }

    public String pixelAfterBucketing(Coordinate looking, Set<Coordinate> filled) {
        if (filled.contains(looking)) return color;
        Set<Coordinate> neighbors = new HashSet<>();
        for(Coordinate p : filled) {
            neighbors.addAll(p.neighboursIn(previousPainter));
        }
        neighbors.removeAll(filled);
        neighbors = neighbors.stream().filter(p-> previousPainter.pixel(p) == " ").collect(Collectors.toSet());
        if(neighbors.size() > 0) {
            filled.addAll(neighbors);
            return pixelAfterBucketing(looking, filled);
        }
        return " ";
    }

}
