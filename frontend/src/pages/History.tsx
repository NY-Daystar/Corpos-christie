import { useEffect, useState } from "react";
import DataGrid from "../components/DataGrid";

import { GetHistory } from "../../wailsjs/go/main/App";

export default function History() {
	const [data, setData] = useState<any[]>([]);

	useEffect(() => {
		GetHistory().then(r => {
			setData(r);
		});
	}, []);

	return (
		<DataGrid
			columns={["Date", "Country", "Gross", "Tax", "Net"]}
			data={data}
		/>
	);
}
