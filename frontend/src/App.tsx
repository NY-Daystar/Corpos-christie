import { useState } from "react";
import Tabs from "./components/Tabs";

import Taxes from "./pages/Taxes";
import History from "./pages/History";
import Settings from "./pages/Settings";

export default function App() {
	const [tab, setTab] = useState("taxes");

	const tabs = [
		{ id: "taxes", label: "Taxes" },
		{ id: "history", label: "History" },
		{ id: "settings", label: "Settings" }
	];

	const render = () => {
		switch (tab) {
			case "taxes":
				return <Taxes />;
			case "history":
				return <History />;
			case "settings":
				return <Settings />;
		}
	};

	return (
		<div>
			<Tabs tabs={tabs} active={tab} onChange={setTab} />

			<div style={{ padding: 20 }}>{render()}</div>
		</div>
	);
}
