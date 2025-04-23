/**
 * 取得中の処理があればその処理を行わずに実行中の処理が完了した結果を受け取る
 *
 * argsをJSON.stringifyした結果が一致していたら同じ結果を利用する
 *
 * 作られた関数は元の結果と共に、ほかの取得の結果を利用したかどうかが帰ってくる
 *
 * ref: https://github.com/traPtitech/traQ_S-UI/blob/cacb18952e97360308e3b7c17f800de3aefefad2/src/lib/basic/async.ts#L8
 */
export const createSingleflight = <T extends unknown[], S>(
  func: (...args: T) => Promise<S>
): ((...args: T) => Promise<[S, boolean]>) => {
  const cacheMap = new Map<string, Promise<S>>()

  return async (...args) => {
    const key = JSON.stringify(args)
    if (cacheMap.has(key)) {
      const promise = cacheMap.get(key)!
      const res = await promise
      return [res, true]
    }

    const promise = func(...args)
    cacheMap.set(key, promise)
    try {
      const res = await promise
      return [res, false]
    } finally {
      cacheMap.delete(key)
    }
  }
}
