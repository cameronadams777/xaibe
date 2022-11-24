export const getElByKey = (obj: Record<any, any>, keys: string[]): any => {
  if (keys.length === 1) return obj[keys[0]];
  return getElByKey(obj[keys[0]], keys.slice(1));
};
