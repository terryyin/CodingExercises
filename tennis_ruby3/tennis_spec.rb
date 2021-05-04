# ruby 3.0.1
#


class TennisGame
  def initialize
    @game_status = [0, 0]
  end

  def score!(player)
    @game_status = rule[player]
    @game_status = [3, 3] if @game_status == [4, 4]
  end

  def score = name % { '0': 'Love', '1': 'Fifteen', '2': 'Thirty', '3': 'Forty', 'p2': 'Player Two', 'p1': 'Player One' }

  private

  def rule
    {player1: [@game_status[0] + 1, @game_status[1]], player2: [@game_status[0], @game_status[1] + 1]}
  end

  def name
    case @game_status
    in [3, 3]
      'Deuce'
    in [a, b] if a == b
      "%{#{a}} All"
    in [3, 4]
      'Advantage %{p2}'
    in [4, 3]
      'Advantage %{p1}'
    in [4..5, _]
      '%{p1} Wins'
    in [_, 4..5]
      '%{p2} Wins'
    else
      "%{#{@game_status[0]}} %{#{@game_status[1]}}"
    end
  end
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

