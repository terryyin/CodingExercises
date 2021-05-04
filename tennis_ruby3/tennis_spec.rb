# ruby 3.0.1
#

class TennisGame
  def initialize
    @score = '0 0'
  end

  RULES = {
      '0 0' => {player1: 'Fifteen Love', player2: 'Love Fifteen'},
      'Fifteen Love' => {player1: 'Thirty Love', player2: 'Fifteen All'},
      'Love Fifteen' => {player1: 'Fifteen All', player2: 'Love Thirty'},
      'Fifteen All' => {player1: 'Thirty Fifteen', player2: 'Fifteen Thirty'},
      'Thirty Fifteen' => { player1: 'Forty Fifteen', player2:'Thirty All'},
      'Fifteen Thirty' => { player1: 'Thirty All', player2:'Fifteen Forty'},
      'Thirty All' => {player1: 'Forty Thirty', player2: 'Thirty Forty'},
      'Forty Thirty' => { player1: 'Player One Wins', player2:'Deuce'},
      'Thirty Forty' => { player1: 'Deuce', player2:'Player Two Wins'},
      'Deuce' => {player1: 'Advantage Player One', player2: 'Advantage Player Two'},
      'Advantage Player One' => {player1: 'Player One Wins', player2: 'Deuce'},
      'Advantage Player Two' => {player1: 'Deuce', player2: 'Player Two Wins'},
  }.freeze

  NAMES = {
    '0 0' => 'Love All',
  }
  def player1_score = @score = RULES[@score][:player1]
  def player2_score = @score = RULES[@score][:player2]

  def score
    return NAMES[@score] if NAMES.keys.include?(@score)
    @score
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
    [5, 3, "Player One Wins"],
    [3, 5, "Player Two Wins"],
    [4, 4, "Deuce"],
    [4, 5, "Advantage Player Two"],
    [4, 6, "Player Two Wins"]
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
end

