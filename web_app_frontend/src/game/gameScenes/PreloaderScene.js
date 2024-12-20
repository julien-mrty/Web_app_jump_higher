import Phaser from "phaser";

//Character
import HeroRunAnimation from "/src/gameAssets/gameSprites/HeroRun.png"
import EnemyRunAttackAnimation from "/src/gameAssets/gameSprites/EnemyRunAttack2.png"

//Environment
import GroundSprite from "/src/gameAssets/gameBackgrounds/tile2.png"
import Fog from "/src/gameAssets/gameBackgrounds/Fog.png"
import Tree1 from "/src/gameAssets/gameBackgrounds/Tree1.png"


//import IdleAnimation from "/src/gameAssets/gameSprites/Idle.png"
//import RunAnimationLeft from "/src/gameAssets/gameSprites/RunReverse.png"


class PreloaderScene extends Phaser.Scene {
    constructor() {
        super("PreloaderScene");
    }

    preload() {
        this.load.image("groundSprite", GroundSprite); //Load a texture that will be named “groundSprite”. Use the same name in MainScence.js to access it
        this.load.image("Fog", Fog); // Load a texture named "Fog"
        this.load.image("tree1", Tree1); // Load a texture named "Tree1"

        // Load a spritesheet with multiple frames
        this.load.spritesheet('RunAnimation', HeroRunAnimation, {
            frameWidth: 42, //width of a single frame in the spridesheet 
            frameHeight: 42 //height of a single frame in the spridesheet
        });//Player sprite

        // Load a spritesheet with multiple frames
        this.load.spritesheet('EnemyRunAttackAnimation', EnemyRunAttackAnimation, {
            frameWidth: 42, //width of a single frame in the spridesheet 
            frameHeight: 42 //height of a single frame in the spridesheet
        });//Player sprite
    }

    create() {
        this.scene.start("MainScene");
    }
}

export default PreloaderScene;