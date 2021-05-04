# ruby 3.0.1
#

class TennisGame
  def initialize
    @score = [0, 0]
  end

  RULES = {
    [0,0] => {name: '%{0} All',},
    [0,1] => {name: 'Love Fifteen',},
    [0,2] => {name: 'Love Thirty',},
    [0,3] => {name: 'Love Forty',},
    [1,0] => {name: 'Fifteen Love',},
    [1,1] => {name: 'Fifteen All',},
    [1,2] => {name: 'Fifteen Thirty',},
    [1,3] => {player2: [0,4], name: 'Fifteen Forty',},
    [2,0] => {name: 'Thirty Love',},
    [2,1] => {name: 'Thirty Fifteen',},
    [2,2] => {name: 'Thirty All',},
    [2,3] => {player2: [0,4], name: 'Thirty Forty',},
    [3,0] => {name: 'Forty Love',},
    [3,1] => {player1: [4,0], name: 'Forty Fifteen',},
    [3,2] => {player1: [4,0], name: 'Forty Thirty',},
    [3,3] => {name: 'Deuce',},
    [3,4] => {player1: [3,3], player2: [0,4], name: 'Advantage Player Two',},
    [4,3] => {player1: [4,0], player2: [3,3], name: 'Advantage Player One',},
    [4,0] => {name: 'Player One Wins',},
    [0,4] => {name: 'Player Two Wins'},
  }.freeze

  def player1_score = @score = RULES[@score].fetch(:player1) {[@score[0] + 1, @score[1]]}
  def player2_score = @score = RULES[@score].fetch(:player2) {[@score[0], @score[1] + 1]}

  def score
    return RULES[@score][:name] % { '0': 'Love' }
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

