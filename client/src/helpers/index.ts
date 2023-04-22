import fromPairs from "lodash/fromPairs";
export * from "./alerts";
export * from "./applications";

export const getSrc = (path: string) => {
  const modules = import.meta.glob(
    "/src/assets/images/**/*.{jpg,png,img,svg}",
    { eager: true }
  );
  return (modules[path] as any)?.default;
};

export const hasTruthyFields = (obj: any): boolean => {
  if (typeof obj !== "object") return false;
  const isUndefOrNull = (val: any) => val == null;
  const stringHasLength = (val: any) =>
    typeof val === "string" ? val.length > 0 : true;

  return Object.values(obj).some(
    (val) => !isUndefOrNull(val) && stringHasLength(val)
  );
};

export const deserializeCookie = (cookieString: string) => {
  const cookieArray = cookieString.split("; ");
  const cookiesObject = fromPairs(cookieArray.map(cookie => cookie.split("="))); 
  return cookiesObject;
}

export const deserializeRawCookies = (cookiesArr: string[]) => {
  const cookiesObject = fromPairs(cookiesArr.map(cookie => cookie.split("="))); 
  return cookiesObject;
}
