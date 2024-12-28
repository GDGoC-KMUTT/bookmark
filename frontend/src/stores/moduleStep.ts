// // import { atom } from "jotai";

// // export const moduleStepsAtom = atom<{
// // 	[moduleId: number]: { step_id: number; title: string; check: string }[];
// //   }>({});

// import { atom } from "jotai";


//   export const moduleStepsAtom = atom<Array<{
// 	step_id: number;
// 	title: string;
// 	check: string;
// }>>([]);

import { atom } from "jotai";

export const moduleStepsAtom = atom<{
  [moduleId: number]: { step_id: number; title: string; check: string }[];
}>({});
