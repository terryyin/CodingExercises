import java.util.ArrayList;
import java.util.List;
import java.util.stream.Collectors;

public class Coordinate {
    public final int x;
    public final int y;

    public Coordinate(int x, int y) {
        this.x = x;
        this.y = y;
    }

    @Override
    public boolean equals(Object o) {
        Coordinate c = (Coordinate) o;
        return x == c.x && y == c.y;
    }

    @Override
    public int hashCode() {
        return x*100 + y;
    }

    public List<Coordinate> neighboursIn(TextPainter textPainter) {
        List<Coordinate> nbs = new ArrayList<Coordinate>();
        nbs.add(new Coordinate(x, y - 1));
        nbs.add(new Coordinate(x, y + 1));
        nbs.add(new Coordinate(x - 1, y));
        nbs.add(new Coordinate(x + 1, y));
        nbs.add(new Coordinate(x - 1, y - 1));
        nbs.add(new Coordinate(x - 1, y + 1));
        nbs.add(new Coordinate(x + 1, y - 1));
        nbs.add(new Coordinate(x + 1, y + 1));
        return nbs.stream().filter(p->p.x > 0 && p.x <= textPainter.canvas.width && p.y > 0 && p.y <= textPainter.canvas.height).collect(Collectors.toList());
    }
}
