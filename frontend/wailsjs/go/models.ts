export namespace models {
	
	export class File {
	    directorypath: string;
	    filename: string;
	    filesize: number;
	    // Go type: time
	    filemodtime: any;
	    depth: number;
	
	    static createFrom(source: any = {}) {
	        return new File(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.directorypath = source["directorypath"];
	        this.filename = source["filename"];
	        this.filesize = source["filesize"];
	        this.filemodtime = this.convertValues(source["filemodtime"], null);
	        this.depth = source["depth"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class ResponseFileStruct {
	    root_path: string;
	    files: {[key: string]: File[]};
	    file: File;
	
	    static createFrom(source: any = {}) {
	        return new ResponseFileStruct(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.root_path = source["root_path"];
	        this.files = source["files"];
	        this.file = this.convertValues(source["file"], File);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

