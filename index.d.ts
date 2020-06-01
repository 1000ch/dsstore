interface RemoveDSStore {
  (): Promise<boolean>;
}

declare module 'dsstore' {
  const f: RemoveDSStore;
  export default f;
}
