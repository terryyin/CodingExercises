require 'spec_helper'

class HuaRongDaoGame
  def self.with(*tiles)
    HuaRongDaoGame.new.add(*tiles)
  end

  def initialize
    @tiles = []
  end

  def solution
    solution_with_path([])
  end

  def solution_with_path(path)
    return nil if path.count > 3
    return [] if @tiles.include? CaoCao.at(1,3)
    [[1, 4], [2, 4], [2, 3]].collect{|x, y| tile_at(x, y)}.compact.
      select{|t| !t.is_a? CaoCao }.each do |tile|
      m = possible_moves(tile)
      if m
        game = HuaRongDaoGame.with(*(@tiles - [tile] + [m]))
        s = game.solution_with_path(path + [self])
        if s
          return [m] + s
        end
      end
    end
    [CaoCao.move_to(1, 3)]
  end

  def add(*tiles)
    @tiles += tiles
    self
  end

  def tile_at(x, y)
    @tiles.select{|t|t.at?(x,y)}.first
  end

  def possible_moves(tile)
    [1, -1].each do |dx|
      m =tile.move(dx, 0)
      return m unless tile_at(m.x, m.y)
    end
    nil
  end
end

class Tile
  attr :x, :y
  def self.at(x, y)
    @@tiles ||= {}
    @@tiles[[self, x, y]] ||= new(x, y)
  end

  def self.move_to(x, y)
    at(x, y)
  end

  def initialize(x, y)
    @x, @y = x, y
  end

  def at?(x, y)
    @x == x && @y == y
  end

  def move_to(x, y)
    self.class.move_to(x, y)
  end

  def move(dx, dy)
    self.class.move_to(@x + dx, @y + dy)
  end
end

class CaoCao < Tile
  def at?(x, y)
    (@x == x || @x==x-1) && (@y == y || @y==y-1)
  end
end

class SoldierA < Tile
end

class SoldierB < Tile
end

class ZhaoYun < Tile
  def at?(x, y)
    @x == x && (@y == y || @y==y-1)
  end
end

describe 'the Hua Rong Dao game' do
  let(:game){Hrd.new()}

  describe "an ended game" do
    let(:game) {HuaRongDaoGame.with(CaoCao.at(1, 3))}
    it { expect(game.solution).to eq [] }
    it { expect(game.add(SoldierA.at(0, 0)).solution).to eq [] }
  end

  describe 'CaoCao need to move to the winning position' do
    it { expect(HuaRongDaoGame.with(CaoCao.at(0, 3)).solution).to eq [CaoCao.move_to(1, 3)]}
    it { expect(HuaRongDaoGame.with(CaoCao.at(2, 3)).solution).to eq [CaoCao.move_to(1, 3)]}
    it { expect(HuaRongDaoGame.with(CaoCao.at(1, 2)).solution).to eq [CaoCao.move_to(1, 3)]}
  end

  describe 'A soldier is blocking Caocao' do
    let(:game) {HuaRongDaoGame.with(CaoCao.at(1, 2))}
    context 'when 34 is taken' do
      before { game.add(SoldierA.at(3,4)) }
      it { expect(game.add(SoldierA.at(1,4)).solution).to eq [SoldierA.move_to(0, 4), CaoCao.move_to(1, 3)]}
      it { expect(game.add(SoldierB.at(1,4)).solution).to eq [SoldierB.move_to(0, 4), CaoCao.move_to(1, 3)]}
      it { expect(game.add(SoldierB.at(2,4)).solution).to include SoldierB.move_to(0, 4)}
    end

    context 'when 04 is taken' do
      before { game.add(SoldierA.at(0,4)) }
      it { expect(game.add(SoldierB.at(1,4)).solution).to include SoldierB.move_to(3, 4)}
    end

    context 'when 03 is taken by ZhaoYun' do
      before { game.add(ZhaoYun.at(0,3)) }
      it { expect(game.add(SoldierB.at(1,4)).solution).to include SoldierB.move_to(3, 4)}
    end

    context 'two soldiers blocking caocao' do
      before { game.add(SoldierA.at(2,4), SoldierB.at(1,4)) }
      it { expect(game.solution).to include SoldierB.move_to(0, 4) }
      it { expect(game.solution).to include SoldierA.move_to(3, 4) }
    end
  end

  describe 'A soldier is at side' do
    let(:game) {HuaRongDaoGame.with(CaoCao.at(0, 3))}
    it { expect(game.add(SoldierA.at(2,3)).solution).to include SoldierA.move_to(3, 3)}

    context 'when 33 is taken' do
      before { game.add(SoldierB.at(3,3)) }
      it { expect(game.add(SoldierA.at(2,3)).solution).to include SoldierA.move_to(3, 4)}
    end
  end
end
