import static org.junit.Assert.*;

import org.junit.Test;


public class TestFizzBuzz {
	FizzBuzz fb = new FizzBuzz();

	@Test
	public void OneShouldBeOne() {
		assertEquals("1", fb.evaluate(1));
	}
	
	@Test
	public void TwoShouldBeTwo() {
		assertEquals("2", fb.evaluate(2));
	}
	
	@Test
	public void ThreeShouldBeFizz() {
		assertEquals("Fizz", fb.evaluate(3));
	}
	
	@Test
	public void SixShouldBeFizzToo() {
		assertEquals("Fizz", fb.evaluate(6));
	}
	
	@Test
	public void FiveShouldBeBuzz() {
		assertEquals("Buzz", fb.evaluate(5));
	}

	@Test
	public void FifteenShouldBeFizzBuzz() {
		assertEquals("FizzBuzz", fb.evaluate(15));
	}
}
