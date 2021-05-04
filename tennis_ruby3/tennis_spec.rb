# ruby 3.0.1
#

class TennisGame
  attr_reader :score

  def initialize
    @score = 'Love All'
  end

  def player1_score
    if @score =='Fifteen All'
      @score = 'Thirty Fifteen'
    elsif @score =='Thirty All'
      @score = 'Forty Thirty'
    else
      @score ='Fifteen Love'
    end
  end

  def player2_score
    if @score == 'Love All'
      @score ='Love Fifteen'
    elsif @score == 'Thirty Fifteen'
      @score ='Thirty All'
    else
      @score ='Fifteen All'
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
    #    [3, 3, "Deuce"],
    #    [4, 3, "Advantage Player One"],
    #    [5, 3, "Player One Wins"],
    #    [3, 5, "Player Two Wins"],
    #    [4, 4, "Deuce"],
    #    [4, 5, "Advantage Player Two"],
    #    [4, 6, "Player Two Wins"]
  ].each do |player1_balls, player2_balls, expectation|
    it "when player1 wins #{player1_balls} and player2 wins #{player2_balls}, should say `#{expectation}`" do
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

  def lead_balls(player1_balls, player2_balls)

  end

end

