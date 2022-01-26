/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "avendauz.tictactoe1.tictactoe1";

export interface OpenGame {
  initiator: string;
  uuid: string;
}

export interface CurrGame {
  initiator: string;
  challenger: string;
  uuid: string;
  board: Uint8Array;
}

export interface DoneGame {
  initiator: string;
  challenger: string;
  uuid: string;
  board: Uint8Array;
  winner: string;
}

const baseOpenGame: object = { initiator: "", uuid: "" };

export const OpenGame = {
  encode(message: OpenGame, writer: Writer = Writer.create()): Writer {
    if (message.initiator !== "") {
      writer.uint32(10).string(message.initiator);
    }
    if (message.uuid !== "") {
      writer.uint32(18).string(message.uuid);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): OpenGame {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseOpenGame } as OpenGame;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.initiator = reader.string();
          break;
        case 2:
          message.uuid = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): OpenGame {
    const message = { ...baseOpenGame } as OpenGame;
    if (object.initiator !== undefined && object.initiator !== null) {
      message.initiator = String(object.initiator);
    } else {
      message.initiator = "";
    }
    if (object.uuid !== undefined && object.uuid !== null) {
      message.uuid = String(object.uuid);
    } else {
      message.uuid = "";
    }
    return message;
  },

  toJSON(message: OpenGame): unknown {
    const obj: any = {};
    message.initiator !== undefined && (obj.initiator = message.initiator);
    message.uuid !== undefined && (obj.uuid = message.uuid);
    return obj;
  },

  fromPartial(object: DeepPartial<OpenGame>): OpenGame {
    const message = { ...baseOpenGame } as OpenGame;
    if (object.initiator !== undefined && object.initiator !== null) {
      message.initiator = object.initiator;
    } else {
      message.initiator = "";
    }
    if (object.uuid !== undefined && object.uuid !== null) {
      message.uuid = object.uuid;
    } else {
      message.uuid = "";
    }
    return message;
  },
};

const baseCurrGame: object = { initiator: "", challenger: "", uuid: "" };

export const CurrGame = {
  encode(message: CurrGame, writer: Writer = Writer.create()): Writer {
    if (message.initiator !== "") {
      writer.uint32(10).string(message.initiator);
    }
    if (message.challenger !== "") {
      writer.uint32(18).string(message.challenger);
    }
    if (message.uuid !== "") {
      writer.uint32(26).string(message.uuid);
    }
    if (message.board.length !== 0) {
      writer.uint32(34).bytes(message.board);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): CurrGame {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseCurrGame } as CurrGame;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.initiator = reader.string();
          break;
        case 2:
          message.challenger = reader.string();
          break;
        case 3:
          message.uuid = reader.string();
          break;
        case 4:
          message.board = reader.bytes();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): CurrGame {
    const message = { ...baseCurrGame } as CurrGame;
    if (object.initiator !== undefined && object.initiator !== null) {
      message.initiator = String(object.initiator);
    } else {
      message.initiator = "";
    }
    if (object.challenger !== undefined && object.challenger !== null) {
      message.challenger = String(object.challenger);
    } else {
      message.challenger = "";
    }
    if (object.uuid !== undefined && object.uuid !== null) {
      message.uuid = String(object.uuid);
    } else {
      message.uuid = "";
    }
    if (object.board !== undefined && object.board !== null) {
      message.board = bytesFromBase64(object.board);
    }
    return message;
  },

  toJSON(message: CurrGame): unknown {
    const obj: any = {};
    message.initiator !== undefined && (obj.initiator = message.initiator);
    message.challenger !== undefined && (obj.challenger = message.challenger);
    message.uuid !== undefined && (obj.uuid = message.uuid);
    message.board !== undefined &&
      (obj.board = base64FromBytes(
        message.board !== undefined ? message.board : new Uint8Array()
      ));
    return obj;
  },

  fromPartial(object: DeepPartial<CurrGame>): CurrGame {
    const message = { ...baseCurrGame } as CurrGame;
    if (object.initiator !== undefined && object.initiator !== null) {
      message.initiator = object.initiator;
    } else {
      message.initiator = "";
    }
    if (object.challenger !== undefined && object.challenger !== null) {
      message.challenger = object.challenger;
    } else {
      message.challenger = "";
    }
    if (object.uuid !== undefined && object.uuid !== null) {
      message.uuid = object.uuid;
    } else {
      message.uuid = "";
    }
    if (object.board !== undefined && object.board !== null) {
      message.board = object.board;
    } else {
      message.board = new Uint8Array();
    }
    return message;
  },
};

const baseDoneGame: object = {
  initiator: "",
  challenger: "",
  uuid: "",
  winner: "",
};

export const DoneGame = {
  encode(message: DoneGame, writer: Writer = Writer.create()): Writer {
    if (message.initiator !== "") {
      writer.uint32(10).string(message.initiator);
    }
    if (message.challenger !== "") {
      writer.uint32(18).string(message.challenger);
    }
    if (message.uuid !== "") {
      writer.uint32(26).string(message.uuid);
    }
    if (message.board.length !== 0) {
      writer.uint32(34).bytes(message.board);
    }
    if (message.winner !== "") {
      writer.uint32(42).string(message.winner);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): DoneGame {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseDoneGame } as DoneGame;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.initiator = reader.string();
          break;
        case 2:
          message.challenger = reader.string();
          break;
        case 3:
          message.uuid = reader.string();
          break;
        case 4:
          message.board = reader.bytes();
          break;
        case 5:
          message.winner = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): DoneGame {
    const message = { ...baseDoneGame } as DoneGame;
    if (object.initiator !== undefined && object.initiator !== null) {
      message.initiator = String(object.initiator);
    } else {
      message.initiator = "";
    }
    if (object.challenger !== undefined && object.challenger !== null) {
      message.challenger = String(object.challenger);
    } else {
      message.challenger = "";
    }
    if (object.uuid !== undefined && object.uuid !== null) {
      message.uuid = String(object.uuid);
    } else {
      message.uuid = "";
    }
    if (object.board !== undefined && object.board !== null) {
      message.board = bytesFromBase64(object.board);
    }
    if (object.winner !== undefined && object.winner !== null) {
      message.winner = String(object.winner);
    } else {
      message.winner = "";
    }
    return message;
  },

  toJSON(message: DoneGame): unknown {
    const obj: any = {};
    message.initiator !== undefined && (obj.initiator = message.initiator);
    message.challenger !== undefined && (obj.challenger = message.challenger);
    message.uuid !== undefined && (obj.uuid = message.uuid);
    message.board !== undefined &&
      (obj.board = base64FromBytes(
        message.board !== undefined ? message.board : new Uint8Array()
      ));
    message.winner !== undefined && (obj.winner = message.winner);
    return obj;
  },

  fromPartial(object: DeepPartial<DoneGame>): DoneGame {
    const message = { ...baseDoneGame } as DoneGame;
    if (object.initiator !== undefined && object.initiator !== null) {
      message.initiator = object.initiator;
    } else {
      message.initiator = "";
    }
    if (object.challenger !== undefined && object.challenger !== null) {
      message.challenger = object.challenger;
    } else {
      message.challenger = "";
    }
    if (object.uuid !== undefined && object.uuid !== null) {
      message.uuid = object.uuid;
    } else {
      message.uuid = "";
    }
    if (object.board !== undefined && object.board !== null) {
      message.board = object.board;
    } else {
      message.board = new Uint8Array();
    }
    if (object.winner !== undefined && object.winner !== null) {
      message.winner = object.winner;
    } else {
      message.winner = "";
    }
    return message;
  },
};

declare var self: any | undefined;
declare var window: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") return globalThis;
  if (typeof self !== "undefined") return self;
  if (typeof window !== "undefined") return window;
  if (typeof global !== "undefined") return global;
  throw "Unable to locate global object";
})();

const atob: (b64: string) => string =
  globalThis.atob ||
  ((b64) => globalThis.Buffer.from(b64, "base64").toString("binary"));
function bytesFromBase64(b64: string): Uint8Array {
  const bin = atob(b64);
  const arr = new Uint8Array(bin.length);
  for (let i = 0; i < bin.length; ++i) {
    arr[i] = bin.charCodeAt(i);
  }
  return arr;
}

const btoa: (bin: string) => string =
  globalThis.btoa ||
  ((bin) => globalThis.Buffer.from(bin, "binary").toString("base64"));
function base64FromBytes(arr: Uint8Array): string {
  const bin: string[] = [];
  for (let i = 0; i < arr.byteLength; ++i) {
    bin.push(String.fromCharCode(arr[i]));
  }
  return btoa(bin.join(""));
}

type Builtin = Date | Function | Uint8Array | string | number | undefined;
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;
