# ruby 3.0.1
#


class TennisGame
  def initialize
    @score = [0, 0]
  end

  NAMES = {
    [0,0] => '%{0} All',
    [1,1] => '%{1} All',
    [2,2] => '%{2} All',
    [3,3] => 'Deuce',
    [3,4] => 'Advantage %{p2}',
    [4,3] => 'Advantage %{p1}',
    [4,0] => '%{p1} Wins',
    [4,1] => '%{p1} Wins',
    [4,2] => '%{p1} Wins',
    [0,4] => '%{p2} Wins',
    [1,4] => '%{p2} Wins',
    [2,4] => '%{p2} Wins',
  }.freeze


  def score!(player) = @score = rule[player]
  def score = name % { '0': 'Love', '1': 'Fifteen', '2': 'Thirty', '3': 'Forty', 'p2': 'Player Two', 'p1': 'Player One' }

  private

  def rule
    case @score
    when [3, 4]
      {player1: [3,3], player2: [0,4]}
    when [4,3]
      {player1: [4,0], player2: [3,3]}
    else
      {player1: [@score[0] + 1, @score[1]], player2: [@score[0], @score[1] + 1]}
    end
  end

  def name = NAMES.fetch(@score, "%{#{@score[0]}} %{#{@score[1]}}")
end

describe TennisGame do
  let(:game) { TennisGame.new }
  [
    [0, 0, "Love All"],
    [1, 0, "Fifteen Love"],
    [0, 1, "Love Fifteen"],
    [1, 1, "Fifteen All"],
    [2, 1, "Thirty Fifteen"],
    [2, 2, "Thirty All"],
    [3, 2, "Forty Thirty"],
    [3, 3, "Deuce"],
    [4, 3, "Advantage Player One"],
    [3, 4, "Advantage Player Two"],
    [4, 2, "Player One Wins"],
    [5, 3, "Player One Wins"],
    [3, 5, "Player Two Wins"],
    [4, 4, "Deuce"],
    [4, 5, "Advantage Player Two"],
    [4, 6, "Player Two Wins"]
  ].each do |player1_balls, player2_balls, expectation|
    it "when player1 wins #{player1_balls} and player2 wins #{player2_balls}, should say `#{expectation}`" do
      (0...[player1_balls, player2_balls].min).each do
        game.score!(:player1)
        game.score!(:player2)
      end
      (0...(player1_balls - player2_balls)).each do
        game.score!(:player1)
      end
      (0...(player2_balls - player1_balls)).each do
        game.score!(:player2)
      end
      expect(game.score).to eq expectation
    end
  end
end

