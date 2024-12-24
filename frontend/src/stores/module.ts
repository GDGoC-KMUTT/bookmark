import { atom } from "jotai";


  export const moduleAtom = atom<Array<{
    module_id: number;
    title: string;
    description: string;
    image_url: string;
}>>([]);

