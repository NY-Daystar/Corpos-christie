import React from "react";

export interface Tab {
	id: string;
	label: string;
}

interface Props {
	tabs: Tab[];
	active: string;
	onChange: (id: string) => void;
}

export default function Tabs({ tabs, active, onChange }: Props) {
	return (
		<div style={container}>
			{tabs.map(t => (
				<button
					key={t.id}
					onClick={() => onChange(t.id)}
					style={{
						...tab,
						borderBottom:
							active === t.id ? "3px solid #0078d4" : "none"
					}}
				>
					{t.label}
				</button>
			))}
		</div>
	);
}

const container: React.CSSProperties = {
	display: "flex",
	gap: 10,
	borderBottom: "1px solid #ccc",
	padding: 10
};

const tab: React.CSSProperties = {
	background: "none",
	border: "none",
	padding: "10px 15px",
	cursor: "pointer",
	fontSize: 16,
	color: "white"
};
