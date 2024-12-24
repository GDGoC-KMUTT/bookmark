import { atom } from "jotai";
import { text } from "stream/consumers";

export const courseName = atom("");

export const courseInfoAtom = atom({
	course_id: 0,
	name: '',
	field: '',
});

export const courseContentAtom = atom<Array<{
	course_id: number;
	order: number;
	type: string;
	text: string;
	module_id: number;
}>>([]);
