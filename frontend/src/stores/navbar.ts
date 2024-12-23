import { atom } from "jotai";
import type { PayloadProfile } from "@/api/api"; // Import the correct type

export const userProfileAtom = atom<PayloadProfile | undefined>(undefined);
export const totalGemsAtom = atom<number | null>(null);
