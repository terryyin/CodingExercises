# This is a solution to the Game Of Life problem without using any Ruby primitive types,
# including nil, bool, number, string, symbal, array, hash, set.
#
# I find this additional constraint useful, and interesting:-)
#

class World
  def initialize(alive_positions=PositionSet.new)
    @all_alives = alive_positions
  end

  def has_life_at!(position)
    @all_alives.adds!(position)
  end

  def has_life_at?(position)
    @all_alives.has?(position)
  end

  def next_world
    World.new(alive_with_two_neighbours + any_with_three_neighbours)
  end

  private

  def alive_with_two_neighbours
    @all_alives.filter { |p| p.all_neighbours.intersect(@all_alives).has_two_positions? }
  end

  def any_with_three_neighbours
    all_related_positions.filter { |n| n.all_neighbours.intersect(@all_alives).has_three_positions? }
  end

  def all_related_positions
    PositionSet.new.tap do |result|
      @all_alives.for_each {|p| result.append(p.all_neighbours)}
    end
  end
end

class Position
  def initialize(x, y)
    @x, @y = x, y
  end

  def x
    @x
  end

  def y
    @y
  end

  def all_neighbours
    PositionSet.new
      .adds!(left)
      .adds!(right)
      .adds!(up)
      .adds!(down)
      .adds!(downleft)
      .adds!(downright)
      .adds!(upleft)
      .adds!(upright)
  end

  def left
    Position.new @x.next, @y
  end

  def right
    Position.new @x.pre, @y
  end

  def up
    Position.new @x, @y.next
  end

  def down
    Position.new @x, @y.pre
  end

  def downleft
    Position.new @x.next, @y.pre
  end

  def upleft
    Position.new @x.next, @y.next
  end

  def upright
    Position.new @x.pre, @y.next
  end

  def downright
    Position.new @x.pre, @y.pre
  end

  def dead_in?(world)
    !alive_in?(world)
  end

  def alive_in?(world)
    world.has_life_at?(self)
  end

  def ==(other)
    x == other.x and y == other.y
  end
end

class Line
  def next
    @next ||= Line.new.set_pre(self)
  end

  def pre
    @pre ||= Line.new.set_next(self)
  end

  def set_next(nex)
    @next = nex
    self
  end

  def set_pre(pre)
    @pre = pre
    self
  end

end

class PositionSet
  def adds!(position)
    @position = position unless @next_holder&.adds!(position)
    @next_holder ||= PositionSet.new
    self
  end
  def has?(position)
    @position == position || @next_holder&.has?(position)
  end

  def for_each(&block)
    yield @position if @position
    @next_holder&.for_each(&block)
  end

  def filter
    PositionSet.new.tap do |result|
      for_each {|p| result.adds!(p) if yield p }
    end
  end

  def intersect(other)
    filter {|p| other.has?(p) }
  end

  def append(other)
    other.for_each {|p| adds!(p) }
    self
  end

  def +(other)
    PositionSet.new.append(self).append(other)
  end

  def has_three_positions?
    @position && @next_holder.has_two_positions?
  end

  def has_two_positions?
    @position && @next_holder.has_one_position?
  end

  def has_one_position?
    @position && @next_holder.empty?
  end

  def empty?
    !@position
  end
end

require 'spec_helper'

describe 'a position' do
  let(:world) {World.new}
  let(:next_world) {world.next_world}
  subject {Position.new(Line.new, Line.new)}
  it {is_expected.to eq subject}
  it {is_expected.not_to eq subject.left}
  it {is_expected.to eq subject.left.right}
  it {is_expected.to eq subject.right.left}
  it {is_expected.not_to eq subject.up}
  it {is_expected.to eq subject.up.down}
  it {expect(subject.up).not_to eq subject.left}
  it {is_expected.to eq subject.up.right.downleft}
  it {is_expected.to eq subject.down.right.upleft}
  it {is_expected.to eq subject.down.left.upright}
  it {is_expected.to eq subject.up.left.downright}
  it {is_expected.to be_dead_in(world)}

  context 'has life' do
    before {world.has_life_at!(subject)}
    it {is_expected.to be_alive_in(world)}
    it {expect(subject.left).to be_dead_in(world)}
    it {is_expected.to be_dead_in(next_world)}

    context 'when another posisition has life as well' do
      before {world.has_life_at!(subject.left)}
      it {is_expected.to be_alive_in(world)}
      it {expect(subject.left).to be_alive_in(world)}
      it {is_expected.to be_dead_in(next_world)}
    end

    context 'with two alive neightbours' do
      before {world.has_life_at!(subject.left)}
      before {world.has_life_at!(subject.right)}
      it {is_expected.to be_alive_in(next_world)}
      context 'with four alive neightbours' do
        before {world.has_life_at!(subject.up)}
        before {world.has_life_at!(subject.down)}
        it {is_expected.to be_dead_in(next_world)}
      end
    end
  end

  context 'has no life' do
    it {is_expected.to be_dead_in(world)}
    context 'with three alive neightbours' do
      before {world.has_life_at!(subject.left)}
      before {world.has_life_at!(subject.right)}
      before {world.has_life_at!(subject.up)}
      it {is_expected.to be_alive_in(next_world)}
    end
  end
end
