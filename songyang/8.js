// 1. 斐波那契

function fib(n) {
  if (n <= 1) return n
  let p = 0,
    q = 1,
    ans = 0

  for (let i = 2; i <= n; i++) {
    ans = (p + q) % 1000000007
    p = q
    q = ans
  }

  return ans
}

// 2. 青蛙跳台阶
function numWays(n) {
  if (n <= 1) return 1

  let p = 1,
    q = 1,
    ans = 0

  for (let i = 0; i < n; i++) {
    ans = (p + q) % 1000000007
    p = q
    q = ans
  }
  return ans
}

// 上面 通过递推 不断更改 n - 1 , n - 2  获取（fn(n - 1) + fn(n - 2)）

// 3. 买卖股票的最佳时机
function maxProfit(prices) {
  const n = prices.length
  let min = prices[0]
  let maxDiff = 0
  for (let i = 1; i < n; i++) {
    maxDiff = Math.max(maxDiff, prices[i] - min)
    min = Math.min(min, prices[i])
  }
  return maxDiff
}

// 记录当前最小值， 每个对比更改最大差值
