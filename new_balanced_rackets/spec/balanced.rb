class String
  def balanced?
    empty? || matches.any?(&method(:balanced_without?))
  end

  private

  def matches
    ['()', '[]', '{}']
  end

  def balanced_without?(brackets)
    return false unless include?(brackets)
    gsub(brackets, '').balanced?
  end
end
