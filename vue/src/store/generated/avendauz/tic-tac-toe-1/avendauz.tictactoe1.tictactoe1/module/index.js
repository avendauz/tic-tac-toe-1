// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgMove } from "./types/tictactoe1/tx";
import { MsgAcceptGame } from "./types/tictactoe1/tx";
import { MsgOpenGame } from "./types/tictactoe1/tx";
const types = [
    ["/avendauz.tictactoe1.tictactoe1.MsgMove", MsgMove],
    ["/avendauz.tictactoe1.tictactoe1.MsgAcceptGame", MsgAcceptGame],
    ["/avendauz.tictactoe1.tictactoe1.MsgOpenGame", MsgOpenGame],
];
export const MissingWalletError = new Error("wallet is required");
export const registry = new Registry(types);
const defaultFee = {
    amount: [],
    gas: "200000",
};
const txClient = async (wallet, { addr: addr } = { addr: "http://localhost:26657" }) => {
    if (!wallet)
        throw MissingWalletError;
    let client;
    if (addr) {
        client = await SigningStargateClient.connectWithSigner(addr, wallet, { registry });
    }
    else {
        client = await SigningStargateClient.offline(wallet, { registry });
    }
    const { address } = (await wallet.getAccounts())[0];
    return {
        signAndBroadcast: (msgs, { fee, memo } = { fee: defaultFee, memo: "" }) => client.signAndBroadcast(address, msgs, fee, memo),
        msgMove: (data) => ({ typeUrl: "/avendauz.tictactoe1.tictactoe1.MsgMove", value: MsgMove.fromPartial(data) }),
        msgAcceptGame: (data) => ({ typeUrl: "/avendauz.tictactoe1.tictactoe1.MsgAcceptGame", value: MsgAcceptGame.fromPartial(data) }),
        msgOpenGame: (data) => ({ typeUrl: "/avendauz.tictactoe1.tictactoe1.MsgOpenGame", value: MsgOpenGame.fromPartial(data) }),
    };
};
const queryClient = async ({ addr: addr } = { addr: "http://localhost:1317" }) => {
    return new Api({ baseUrl: addr });
};
export { txClient, queryClient, };
