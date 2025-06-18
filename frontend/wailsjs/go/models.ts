export namespace libmanager {
	
	export class Track {
	    id: string;
	    filePath: string;
	    title: string;
	    format: string;
	    album: string;
	    artist: string;
	    original: string;
	    // Go type: time
	    dateAdded: any;
	
	    static createFrom(source: any = {}) {
	        return new Track(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.filePath = source["filePath"];
	        this.title = source["title"];
	        this.format = source["format"];
	        this.album = source["album"];
	        this.artist = source["artist"];
	        this.original = source["original"];
	        this.dateAdded = this.convertValues(source["dateAdded"], null);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
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

