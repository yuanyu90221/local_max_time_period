export interface ReachDataFormat {
  diff: number,
  min: number,
  max: number
}
export interface InputArgs {
  history: Array<number>,
  targetIdx: number
}
export const FindBestInOut = (history: Array<number>):ReachDataFormat => {
  const result = {diff:0, min:0, max: 0};
  const len = history.length;
  for (let idx = 0 ; idx < len; idx++) {
    let tempResult = FindCurrentBest({history, targetIdx: idx});
    if (result.diff <= tempResult.diff) {
      result.diff = tempResult.diff
      result.max = tempResult.max;
      result.min = tempResult.min;
    }
  }
  return result;
} 

const FindCurrentBest = (input:InputArgs):ReachDataFormat => {
  const result = {diff:0, min:0, max: 0};
  const {history, targetIdx} = input;
  const len = history.length;
  for (let idx = 0; idx < len; idx++) {
    if (targetIdx === idx )
      continue;
    let diff = Diff(history[idx], history[targetIdx])
    let max = Max(idx, targetIdx);
    let min = Min(idx, targetIdx);
    if (result.diff <= diff && history[max] > history[min]) {
      result.diff = diff;
      if (max === idx) {
        result.max = idx;
        result.min = targetIdx;
      } else if (idx < targetIdx) {
        result.max = targetIdx;
        result.min = idx;
      }
    }
  }
  return result;
}

const Max = (a:number, b:number): number => {
  return a>b? a:b;
}

const Min = (a:number, b:number): number => {
  return a<b? a:b;
}

const Diff = (a:number, b:number): number => {
  return a > b ? a-b: b-a;
}