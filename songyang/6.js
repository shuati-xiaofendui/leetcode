/**
 * 从上到下打印二叉树
 * https://leetcode.cn/problems/cong-shang-dao-xia-da-yin-er-cha-shu-lcof/
 */

function levelOrder(root) {
  if (!root) return []
  const res = []
  let q = [root]

  while (q.length) {
    const temp = []

    for (let item of q) {
      const [val, left, right] = item

      res.push(val)

      if (left) temp.push(left)
      if (right) temp.push(right)
    }
    q = temp
  }
  return res
}

function levelOrder2(root) {
  if (!root) return []
  const res = []
  let q = [root]

  while (q.length) {
    const temp = []
    const value = []

    for (let item of q) {
      const [val, left, right] = item

      value.push(val)
      if (left) temp.push(left)
      if (right) temp.push(right)
    }
    q = temp
    res.push(value)
  }
  return res
}

function levelOrder3(root) {
  if (!root) return []
  const res = []
  let q = [root]
  let isStartLeft = true

  while (q.length) {
    const temp = []
    const value = []

    for (let item of q) {
      const [val, left, right] = item
      if (isStartLeft) {
        value.push(val)
      } else {
        value.unshift(val)
      }
      if (left) temp.push(left)
      if (right) temp.push(right)
    }
    q = temp
    res.push(value)
  }
  return res
}
