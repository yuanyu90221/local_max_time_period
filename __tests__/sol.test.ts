import { FindBestInOut } from "../src/sol"

test('test_find_local_max_time',(done)=>{
  const history = [[60,100, 58, 20, 56, 42, 20, 250, 194, 10, 1000],[60, 100, 58, 20, 56, 42, 40, 250, 194, 1000, 10]]
  const values = [{diff:990, max:10, min: 9}, {diff: 980, max:9, min: 3}]
  const len = history.length;
  for (let idx = 0; idx < len; idx++) {
    const result = FindBestInOut(history[idx]);
    expect(result).toStrictEqual(values[idx]) 
  }
  done()
})