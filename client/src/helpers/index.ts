export * from "./applications";

export const getSrc = (path: string) => {
  const modules = import.meta.glob(
    "/src/assets/images/**/*.{jpg,png,img,svg}",
    { eager: true }
  );
  return (modules[path] as any)?.default;
};
