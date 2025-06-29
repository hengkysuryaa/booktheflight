import SeatSelection from './SeatSelection';

const App = () => {
  return (
    <div>
      <header className="header">
        <h1>Book The Flight</h1>
        <p>Select your preferred seat from the map below</p>
      </header>
      <main>
        <SeatSelection />
      </main>
    </div>
  );
};

export default App;
