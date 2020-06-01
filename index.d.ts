interface RemoveDSStore {
  (targetPath: string): Promise<boolean>;
}

declare module 'dsstore' {
  const f: RemoveDSStore;
  export default f;
}
