// TODO voir pour les styles comme en react native

import { useState, useEffect } from "react";
import { GetDefaultYear, GetYears } from "../../wailsjs/go/main/App";

export default function Taxes() {
	const options: number[] = Array.from(Array(10).keys());
	const [years, setYears] = useState<string[]>([]);

	useEffect(() => {
		GetYears().then(setYears);
	}, []);

	return (
		<div style={grid}>
			<div style={form}>
				<div>
					<span style={{ marginRight: "15%" }}>
						Enter your income
					</span>
					<input style={{ width: "25%" }} type="text" name="amount" />
				</div>

				<div>
					<span style={{ marginRight: "15%" }}>Statut marital</span>
					<label>
						<input type="radio" name="tax" />
						Single
					</label>
					<label>
						<input type="radio" name="tax" />
						Maried
					</label>
				</div>

				<div>
					<span style={{ marginRight: "25%" }}>Nombre d'enfant</span>
					<select style={{ width: "50px" }}>
						{options.map(opt => (
							<option>{opt}</option>
						))}
					</select>
				</div>
			</div>

			<div>
				<div style={{ display: "flex", justifyContent: "flex-end" }}>
					<select>
						{years.map(year => (
							<option>{year}</option>
						))}
					</select>
				</div>

				<div style={{ marginTop: 30 }}>
					<div style={{ display: "flex", gap: 20 }}>
						<label>Gross</label>
						<label>Tax</label>
						<label>Net</label>
					</div>

					<table style={{ width: "100%", marginTop: 10 }}>
						<thead>
							<tr>
								<th>A</th>
								<th>B</th>
								<th>C</th>
								<th>D</th>
								<th>E</th>
							</tr>
						</thead>

						<tbody>
							{Array.from({ length: 5 }).map((_, i) => (
								<tr key={i}>
									<td></td>
									<td></td>
									<td></td>
									<td></td>
									<td></td>
								</tr>
							))}
						</tbody>
					</table>
				</div>
			</div>
		</div>
	);
}

const grid: React.CSSProperties = {
	display: "grid",
	gridTemplateColumns: "1fr 1fr",
	gap: 40
};

const form: React.CSSProperties = {
	display: "flex",
	flexDirection: "column",
	gap: "2em"
};
