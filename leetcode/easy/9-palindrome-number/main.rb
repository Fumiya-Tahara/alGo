# @param {Integer} x
# @return {Boolean}
def is_palindrome(x)
    if x < 0
        return false
    end

    original = x
    reverse = 0

    while x > 0
        reverse = 10 * reverse + x % 10
        x = x / 10
    end

    if original == reverse
        return true
    end

    return false
end
