// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT
import {models} from '../models';

export function GetPort():Promise<number>;

export function ReceiveMessages():Promise<void>;

export function SendDirectoryContent(arg1:models.ResponseFileStruct):Promise<void>;

export function SendFile(arg1:string,arg2:string):Promise<void>;

export function SetServerPort(arg1:number):Promise<boolean>;
