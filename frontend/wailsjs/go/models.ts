export namespace model {
	
	export class History {
	    date: string;
	    income: number;
	    couple: boolean;
	    children: number;
	
	    static createFrom(source: any = {}) {
	        return new History(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.date = source["date"];
	        this.income = source["income"];
	        this.couple = source["couple"];
	        this.children = source["children"];
	    }
	}

}

