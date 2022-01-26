/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";

export const protobufPackage = "avendauz.tictactoe1.tictactoe1";

export interface MsgOpenGame {
  creator: string;
  uuid: string;
}

export interface MsgOpenGameResponse {}

export interface MsgAcceptGame {
  creator: string;
  uuid: string;
  challenger: string;
}

export interface MsgAcceptGameResponse {}

const baseMsgOpenGame: object = { creator: "", uuid: "" };

export const MsgOpenGame = {
  encode(message: MsgOpenGame, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.uuid !== "") {
      writer.uint32(18).string(message.uuid);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgOpenGame {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgOpenGame } as MsgOpenGame;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
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

  fromJSON(object: any): MsgOpenGame {
    const message = { ...baseMsgOpenGame } as MsgOpenGame;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.uuid !== undefined && object.uuid !== null) {
      message.uuid = String(object.uuid);
    } else {
      message.uuid = "";
    }
    return message;
  },

  toJSON(message: MsgOpenGame): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.uuid !== undefined && (obj.uuid = message.uuid);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgOpenGame>): MsgOpenGame {
    const message = { ...baseMsgOpenGame } as MsgOpenGame;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.uuid !== undefined && object.uuid !== null) {
      message.uuid = object.uuid;
    } else {
      message.uuid = "";
    }
    return message;
  },
};

const baseMsgOpenGameResponse: object = {};

export const MsgOpenGameResponse = {
  encode(_: MsgOpenGameResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgOpenGameResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgOpenGameResponse } as MsgOpenGameResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgOpenGameResponse {
    const message = { ...baseMsgOpenGameResponse } as MsgOpenGameResponse;
    return message;
  },

  toJSON(_: MsgOpenGameResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<MsgOpenGameResponse>): MsgOpenGameResponse {
    const message = { ...baseMsgOpenGameResponse } as MsgOpenGameResponse;
    return message;
  },
};

const baseMsgAcceptGame: object = { creator: "", uuid: "", challenger: "" };

export const MsgAcceptGame = {
  encode(message: MsgAcceptGame, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.uuid !== "") {
      writer.uint32(18).string(message.uuid);
    }
    if (message.challenger !== "") {
      writer.uint32(26).string(message.challenger);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgAcceptGame {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgAcceptGame } as MsgAcceptGame;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.uuid = reader.string();
          break;
        case 3:
          message.challenger = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgAcceptGame {
    const message = { ...baseMsgAcceptGame } as MsgAcceptGame;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.uuid !== undefined && object.uuid !== null) {
      message.uuid = String(object.uuid);
    } else {
      message.uuid = "";
    }
    if (object.challenger !== undefined && object.challenger !== null) {
      message.challenger = String(object.challenger);
    } else {
      message.challenger = "";
    }
    return message;
  },

  toJSON(message: MsgAcceptGame): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.uuid !== undefined && (obj.uuid = message.uuid);
    message.challenger !== undefined && (obj.challenger = message.challenger);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgAcceptGame>): MsgAcceptGame {
    const message = { ...baseMsgAcceptGame } as MsgAcceptGame;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.uuid !== undefined && object.uuid !== null) {
      message.uuid = object.uuid;
    } else {
      message.uuid = "";
    }
    if (object.challenger !== undefined && object.challenger !== null) {
      message.challenger = object.challenger;
    } else {
      message.challenger = "";
    }
    return message;
  },
};

const baseMsgAcceptGameResponse: object = {};

export const MsgAcceptGameResponse = {
  encode(_: MsgAcceptGameResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgAcceptGameResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgAcceptGameResponse } as MsgAcceptGameResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgAcceptGameResponse {
    const message = { ...baseMsgAcceptGameResponse } as MsgAcceptGameResponse;
    return message;
  },

  toJSON(_: MsgAcceptGameResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<MsgAcceptGameResponse>): MsgAcceptGameResponse {
    const message = { ...baseMsgAcceptGameResponse } as MsgAcceptGameResponse;
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  /** this line is used by starport scaffolding # proto/tx/rpc */
  OpenGame(request: MsgOpenGame): Promise<MsgOpenGameResponse>;
  AcceptGame(request: MsgAcceptGame): Promise<MsgAcceptGameResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  OpenGame(request: MsgOpenGame): Promise<MsgOpenGameResponse> {
    const data = MsgOpenGame.encode(request).finish();
    const promise = this.rpc.request(
      "avendauz.tictactoe1.tictactoe1.Msg",
      "OpenGame",
      data
    );
    return promise.then((data) => MsgOpenGameResponse.decode(new Reader(data)));
  }

  AcceptGame(request: MsgAcceptGame): Promise<MsgAcceptGameResponse> {
    const data = MsgAcceptGame.encode(request).finish();
    const promise = this.rpc.request(
      "avendauz.tictactoe1.tictactoe1.Msg",
      "AcceptGame",
      data
    );
    return promise.then((data) =>
      MsgAcceptGameResponse.decode(new Reader(data))
    );
  }
}

interface Rpc {
  request(
    service: string,
    method: string,
    data: Uint8Array
  ): Promise<Uint8Array>;
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
