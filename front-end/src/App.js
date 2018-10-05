import React, { Component } from 'react';
import './App.css';
import 'tachyons';
import SongQueue from './components/SongQueue/SongQueue';

const serverUrl = '192.168.2.14';
const port = '8000';

class App extends Component {
	constructor() {
		super()
		this.state = {
			songUrl: '',
			searchResults: [{}]
		}
	}

	async performSearch() {
		let data = await fetch(`http://${serverUrl}:${port}/search?query=${this.state.songUrl}`)
			.then(res => res.json())
		this.setState({
			songUrl: '',
			searchResults: data
		})

		document.getElementById('search-box').value = '';
	}

	onSkipPressed = () => {
		fetch(`http://${serverUrl}:${port}/skip`)
	}

	onPlayPressed = () => {
		fetch(`http://${serverUrl}:${port}/play?url=${this.state.songUrl}`, {
			method: 'post'
		})
		this.setState({
			songUrl: ''
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

	onSongClicked = (obj) => {
		this.setState({
			songUrl: obj.url,
			searchResults: [{}]
		}, () => this.onPlayPressed())
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
						<button id="add" onClick={this.onPlayPressed} className="mr3 ml3"> + </button>
						<button id="pause" onClick={this.onStopPressed} className="mr3 ml3"> pause </button>
					</div>
				</div>
				<div className='mt3'>
					{
						this.state.searchResults.length === 1 ?
							null : (
								<SongQueue
									searchResults={this.state.searchResults}
									className='flex'
									onSongClicked={this.onSongClicked} />
							)
					}
				</div>
			</div>
		);
	}
}

export default App;
