import React, { Component } from 'react';
import './App.css';
import 'tachyons';
import SongQueue from './components/SongQueue/SongQueue';

class App extends Component {
	constructor() {
		super()
		this.state = {
			songUrl: '',
			searchResults: [{}]
		}
	}

	async performSearch() {
		let data = await fetch('http://192.168.2.14:1100')
			.then(res => res.json())
		this.setState({
			searchResults: data
		})
	}

	onSkipPressed = () => {
		fetch("http://192.168.2.14:8000/skip")
	}

	onPlayPressed = () => {
		fetch(`http://192.168.2.14:8000/play?url=${this.state.songUrl}`, {
			method: 'post'
		})

		document.getElementById('search-box').value = '';
	}

	onStopPressed = () => {
		console.log('Stop pressed')
	}

	onUrlChanged = (event) => {
		this.setState({
			songUrl: event.target.value
		})
	}

	render() {
		return (
			<div className="App">
				<div className="pt3 measure center">
					<div className="ma2">
						<input id="search-box" type="text" name="songUrl" onChange={this.onUrlChanged} />
						<button id="search" onClick={() => this.performSearch()} className="mt1 ml2"> Search </button>
					</div>
					<div className="mt2 pl2 pr2">
						<button id="skip" onClick={this.onSkipPressed} className="mr3 ml3"> skip </button>
						<button id="play" onClick={this.onPlayPressed} className="mr3 ml3"> play </button>
						<button id="stop" onClick={this.onStopPressed} className="mr3 ml3"> stop </button>
					</div>
				</div>
				<div>
					<SongQueue searchResults={this.state.searchResults} />
				</div>
			</div>
		);
	}
}

export default App;
