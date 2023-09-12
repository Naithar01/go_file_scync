export namespace main {
	
	export class ResponseFileStruct {
	    root_path: string;
	    files: {[key: string]: File[]};
	
	    static createFrom(source: any = {}) {
	        return new ResponseFileStruct(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.root_path = source["root_path"];
	        this.files = source["files"];
	    }
	}

}

