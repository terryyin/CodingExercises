# ruby 3.0.1
#

class TennisGame
  attr_reader :score

  def initialize
    @score = 'Love All'
  end

  def player1_score
    @score ='Fifteen Love'
  end

  def player2_score
  end
end

describe TennisGame do
  [
    [0, 0, "Love All"],
    [1, 0, "Fifteen Love"],
#    [1, 1, "Fifteen All"],
#    [2, 1, "Thirty Fifteen"],
#    [2, 2, "Thirty All"],
#    [3, 2, "Forty Thirty"],
#    [3, 3, "Deuce"],
#    [4, 3, "Advantage Player One"],
#    [5, 3, "Player One Wins"],
#    [3, 5, "Player Two Wins"],
#    [4, 4, "Deuce"],
#    [4, 5, "Advantage Player Two"],
#    [4, 6, "Player Two Wins"]
  ].each do |player1_balls, player2_balls, expectation|
    it "when player1 wins #{player1_balls} and player2 wins #{player2_balls}, should say `#{expectation}`" do
      game = TennisGame.new
      (0...[player1_balls, player2_balls].min).each do
        game.player1_score
        game.player2_score
      end
      (0...(player1_balls - player2_balls)).each do
        game.player1_score
      end
      (0...(player2_balls - player1_balls)).each do
        game.player2_score
      end
      expect(game.score).to eq expectation
    end
  end

end

