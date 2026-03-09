import { useState } from "react";

interface Props {
	columns: string[];
	data: any[];
}

export default function DataGrid({ columns, data }: Props) {
	const [page, setPage] = useState(0);
	const pageSize = 5;

	const start = page * pageSize;
	const rows = data.slice(start, start + pageSize);

	return (
		<div>
			<table style={table}>
				<thead>
					<tr>
						{columns.map(c => (
							<th key={c}>{c}</th>
						))}
					</tr>
				</thead>

				<tbody>
					{rows.map((r, i) => (
						<tr key={i}>
							{columns.map(c => (
								<td key={c}>{r[c]}</td>
							))}
						</tr>
					))}
				</tbody>
			</table>

			<div style={{ marginTop: 10 }}>
				<button
					disabled={page === 0}
					onClick={() => setPage(p => p - 1)}
				>
					Prev
				</button>

				<button
					disabled={(page + 1) * pageSize >= data.length}
					onClick={() => setPage(p => p + 1)}
				>
					Next
				</button>
			</div>
		</div>
	);
}

const table: React.CSSProperties = {
	width: "100%",
	borderCollapse: "collapse"
};
