import { Writer, Reader } from "protobufjs/minimal";
export declare const protobufPackage = "avendauz.tictactoe1.tictactoe1";
export interface OpenGame {
    initiator: string;
    uuid: string;
}
export interface CurrGame {
    x: string;
    o: string;
    uuid: string;
    board: Uint8Array;
    turn: string;
}
export interface DoneGame {
    initiator: string;
    challenger: string;
    uuid: string;
    board: Uint8Array;
    winner: string;
}
export declare const OpenGame: {
    encode(message: OpenGame, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): OpenGame;
    fromJSON(object: any): OpenGame;
    toJSON(message: OpenGame): unknown;
    fromPartial(object: DeepPartial<OpenGame>): OpenGame;
};
export declare const CurrGame: {
    encode(message: CurrGame, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): CurrGame;
    fromJSON(object: any): CurrGame;
    toJSON(message: CurrGame): unknown;
    fromPartial(object: DeepPartial<CurrGame>): CurrGame;
};
export declare const DoneGame: {
    encode(message: DoneGame, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): DoneGame;
    fromJSON(object: any): DoneGame;
    toJSON(message: DoneGame): unknown;
    fromPartial(object: DeepPartial<DoneGame>): DoneGame;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
