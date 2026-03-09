import { useEffect, useState } from "react";
import { GetCurrencies, GetDefaultCurrency } from "../../wailsjs/go/main/App";

export default function Settings() {
	const [currencies, setCurrencies] = useState<string[]>([]);
	const [currency, setCurrency] = useState("");

	useEffect(() => {
		GetCurrencies().then(setCurrencies);

		GetDefaultCurrency().then(c => {
			if (c) setCurrency(c);
		});
	}, []);

	return (
		<div>
			<label style={{ marginRight: "10%" }}>Currency</label>

			<select
				value={currency}
				onChange={e => setCurrency(e.target.value)}
			>
				{currencies.map(c => (
					<option key={c} value={c}>
						{c}
					</option>
				))}
			</select>
		</div>
	);
}
