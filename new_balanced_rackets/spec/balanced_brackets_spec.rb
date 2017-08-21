require 'spec_helper'
require 'balanced.rb'

describe 'BalancedBrackets' do
  it { expect('').to be_balanced }
  it { expect('(').not_to be_balanced }
  it { expect('()').to be_balanced }
  it { expect('()()').to be_balanced }
  it { expect('())').not_to be_balanced }
  it { expect('(())').to be_balanced }
  it { expect('((()))').to be_balanced }
  it { expect('((()))(').not_to be_balanced }
  it { expect('[]').to be_balanced }
  it { expect('{}').to be_balanced }
  it { expect('[()]').to be_balanced }
  it { expect('([])').to be_balanced }
end
