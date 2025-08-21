import { writable } from "svelte/store"

export const arrayWritable = <Data = unknown>(initialValue: Data[] = []) => {
  const { subscribe, update, set } = writable<Data[]>(initialValue)

  return {
    subscribe,
    update,
    deleteFromIndex: (...indices: number[]) => update((prevData) => {
      return prevData.filter((_, index) => !indices.includes(index))
    }),
    clear: () => set([])
  }
}