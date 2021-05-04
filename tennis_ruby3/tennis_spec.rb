# ruby 3.0.1
#

class TennisGame
  def initialize
    @score = '0 0'
  end

  RULES = {
    '0 0' => {player1: '1 0', player2: '0 1'},
    '1 0' => {player1: '2 0', player2: '1 1'},
    '0 1' => {player1: '1 1', player2: '0 2'},
    '2 0' => {player1: '3 0', player2: '2 1'},
    '0 2' => {player1: '1 2', player2: '0 3'},
    '1 1' => {player1: '2 1', player2: '1 2'},
    '2 1' => {player1: '3 1', player2:'2 2'},
    '1 2' => {player1: '2 2', player2:'1 3'},
    '3 1' => {player1: ' 1 ', player2: '3 2'},
    '2 2' => {player1: '3 2', player2: '2 3'},
    '3 2' => {player1: ' 1 ', player2:'3 3'},
    '2 3' => {player1: '3 3', player2:' 2 '},
    '3 3' => {player1: '4 3', player2: '3 4'},
    '4 3' => {player1: ' 1 ', player2: '3 3'},
    '3 4' => {player1: '3 3', player2: ' 2 '},
  }.freeze

  NAMES = {
    '0 0' => 'Love All',
    '0 1' => 'Love Fifteen',
    '0 2' => 'Love Thirty',
    '1 0' => 'Fifteen Love',
    '1 1' => 'Fifteen All',
    '1 2' => 'Fifteen Thirty',
    '1 3' => 'Fifteen Forty',
    '2 0' => 'Thirty Love',
    '2 1' => 'Thirty Fifteen',
    '2 2' => 'Thirty All',
    '2 3' => 'Thirty Forty',
    '3 1' => 'Forty Fifteen',
    '3 2' => 'Forty Thirty',
    '3 3' => 'Deuce',
    '3 4' => 'Advantage Player Two',
    '4 3' => 'Advantage Player One',
    ' 1 ' => 'Player One Wins',
    ' 2 ' => 'Player Two Wins'
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

