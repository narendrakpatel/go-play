import React from 'react';
import './Song.css'

const Song = ({title, url, thumb}) => {
    return (
        <div className='tc dib br3 pa3 ma2 grow bw2 shadow-5'>
            <img src={thumb} alt={title} className='dimention'/><br/>
            <label className='label'>{title}</label><br/>
            <label className='label'>{url}</label>
        </div>
    )
}

export default Song;