import org.junit.Assert;
import org.junit.Ignore;
import org.junit.Test;

public class TestTextPainter {
    TextPainter tp = new TextPainter(20, 4);

    @Test
    public void HorizontalLine() {
        tp = tp.command("L 1 2 6 2");
        Assert.assertEquals(" ", tp.pixel(1, 1));
        Assert.assertEquals("X", tp.pixel(1, 2));
        Assert.assertEquals("X", tp.pixel(2, 2));
        Assert.assertEquals(" ", tp.pixel(7, 2));
    }

    @Test
    public void VerticalLine() {
        tp = tp.command("L 6 3 6 4");
        Assert.assertEquals(" ", tp.pixel(6, 2));
        Assert.assertEquals("X", tp.pixel(6, 4));
        Assert.assertEquals("X", tp.pixel(6, 3));
        Assert.assertEquals(" ", tp.pixel(5, 3));
        Assert.assertEquals(" ", tp.pixel(2, 2));
    }

    @Test
    public void HorizonTalAndVerticalLine() {
        tp = tp.command("L 1 2 6 2");
        tp = tp.command("L 6 3 6 4");
        Assert.assertEquals("X", tp.pixel(2, 2));
        Assert.assertEquals("X", tp.pixel(6, 3));
    }

    @Test
    public void Rectangle() {
        tp = tp.command("R 14 1 18 3");
        Assert.assertEquals(" ", tp.pixel(15, 2));
        Assert.assertEquals("X", tp.pixel(15, 1));
        Assert.assertEquals("X", tp.pixel(14, 2));
        Assert.assertEquals("X", tp.pixel(18, 2));
        Assert.assertEquals("X", tp.pixel(15, 3));
    }

    @Test
    public void BucketAnEmptyCanvas() {
        tp = tp.command("B 10 3 o");
        Assert.assertEquals("o", tp.pixel(10, 3));
    }

    @Test
    public void BucketCanvasWithALine() {
        tp = tp.command("L 1 2 6 2");
        tp = tp.command("B 10 3 o");
        Assert.assertEquals("X", tp.pixel(1, 2));
    }

    @Test
    public void BucketCanvasWithABox() {
        tp = tp.command("R 14 1 18 3");
        tp = tp.command("B 10 3 o");
        Assert.assertEquals("o", tp.pixel(10, 3));
        Assert.assertEquals("o", tp.pixel(10, 2));
        Assert.assertEquals("o", tp.pixel(10, 1));
        Assert.assertEquals("o", tp.pixel(10, 4));
        Assert.assertEquals("o", tp.pixel(11, 4));
        Assert.assertEquals(" ", tp.pixel(15, 2));
    }
}
