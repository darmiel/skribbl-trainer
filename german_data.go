package main

const GermanDataJson = `
{
  "Sommersprossen": {
    "count": 5,
    "lastSeenTime": 1586959168567
  },
  "inkognito": {
    "count": 7,
    "lastSeenTime": 1586901586361
  },
  "Marienkäfer": {
    "count": 7,
    "lastSeenTime": 1586958264090,
    "difficulty": 0.8902439024390244,
    "difficultyWeight": 82
  },
  "Zaun": {
    "count": 8,
    "lastSeenTime": 1586958848320
  },
  "graben": {
    "count": 6,
    "lastSeenTime": 1586896977731,
    "publicGameCount": 1
  },
  "Erfrierung": {
    "count": 7,
    "lastSeenTime": 1586900635930
  },
  "Kessel": {
    "count": 7,
    "lastSeenTime": 1586839270292
  },
  "Kellner": {
    "count": 6,
    "lastSeenTime": 1586900763534
  },
  "Panzerband": {
    "count": 4,
    "lastSeenTime": 1586860197013,
    "difficulty": 1,
    "difficultyWeight": 11
  },
  "Schnorchel": {
    "count": 12,
    "lastSeenTime": 1586901554174,
    "difficulty": 1,
    "difficultyWeight": 4
  },
  "Abgrund": {
    "count": 10,
    "lastSeenTime": 1586869078737,
    "difficulty": 1,
    "difficultyWeight": 2
  },
  "geduldig": {
    "count": 8,
    "lastSeenTime": 1586959178759,
    "publicGameCount": 1,
    "difficulty": 0.8888888888888888,
    "difficultyWeight": 9
  },
  "Edelstein": {
    "count": 8,
    "lastSeenTime": 1586954627819,
    "difficulty": 1,
    "difficultyWeight": 1
  },
  "Rasseln": {
    "count": 6,
    "lastSeenTime": 1586919552513,
    "difficulty": 1,
    "difficultyWeight": 3
  },
  "Ohr": {
    "count": 11,
    "lastSeenTime": 1586900881142,
    "difficulty": 0.5833333333333334,
    "difficultyWeight": 12
  },
  "Elektron": {
    "count": 9,
    "lastSeenTime": 1586903684169
  },
  "Telefonbuch": {
    "count": 7,
    "lastSeenTime": 1586908504643
  },
  "Kim Jong-Un": {
    "count": 5,
    "lastSeenTime": 1586904405029
  },
  "Backenzahn": {
    "count": 9,
    "lastSeenTime": 1586915737415
  },
  "Ameisenbär": {
    "count": 9,
    "lastSeenTime": 1586916550509,
    "difficulty": 0.8,
    "difficultyWeight": 5
  },
  "Zaubertrick": {
    "count": 4,
    "lastSeenTime": 1586783316765,
    "difficulty": 1,
    "difficultyWeight": 4
  },
  "Angestellter": {
    "count": 5,
    "lastSeenTime": 1586858975693,
    "publicGameCount": 1
  },
  "Beförderung": {
    "count": 6,
    "lastSeenTime": 1586914119532,
    "difficulty": 1,
    "difficultyWeight": 6
  },
  "Krokodil": {
    "count": 8,
    "lastSeenTime": 1586901461723
  },
  "keuchen": {
    "count": 6,
    "lastSeenTime": 1586875186652
  },
  "Bleiche": {
    "count": 6,
    "lastSeenTime": 1586785407694
  },
  "The Rock": {
    "count": 8,
    "lastSeenTime": 1586901554174
  },
  "Karteikasten": {
    "count": 8,
    "lastSeenTime": 1586958934411,
    "difficulty": 0.875,
    "difficultyWeight": 8
  },
  "fliegendes Schwein": {
    "count": 13,
    "lastSeenTime": 1586920922732,
    "difficulty": 0.9473684210526315,
    "difficultyWeight": 19
  },
  "wachsen": {
    "count": 7,
    "lastSeenTime": 1586920637722,
    "difficulty": 1,
    "difficultyWeight": 1
  },
  "scharfe Soße": {
    "count": 4,
    "lastSeenTime": 1586959595986
  },
  "Grab": {
    "count": 8,
    "lastSeenTime": 1586919297771
  },
  "Milchmann": {
    "count": 10,
    "lastSeenTime": 1586955716758
  },
  "Pfanne": {
    "count": 11,
    "lastSeenTime": 1586958959180
  },
  "Säbel": {
    "count": 11,
    "lastSeenTime": 1586900301076,
    "difficulty": 0.75,
    "difficultyWeight": 28
  },
  "Peperoni": {
    "count": 9,
    "lastSeenTime": 1586901643471,
    "difficulty": 1,
    "difficultyWeight": 17
  },
  "Pony": {
    "count": 6,
    "lastSeenTime": 1586914196223
  },
  "Zirkus": {
    "count": 8,
    "lastSeenTime": 1586959814583,
    "publicGameCount": 1,
    "difficulty": 0.375,
    "difficultyWeight": 8
  },
  "Kredithai": {
    "count": 2,
    "lastSeenTime": 1586784756252,
    "publicGameCount": 1,
    "difficulty": 0.8,
    "difficultyWeight": 5
  },
  "Triangel": {
    "count": 4,
    "lastSeenTime": 1586954834961
  },
  "Jo-Jo": {
    "count": 10,
    "lastSeenTime": 1586919312988,
    "publicGameCount": 1
  },
  "Tür": {
    "count": 8,
    "lastSeenTime": 1586826993481,
    "difficulty": 0.8181818181818182,
    "difficultyWeight": 11
  },
  "Bauchnabel": {
    "count": 6,
    "lastSeenTime": 1586896963540,
    "difficulty": 1,
    "difficultyWeight": 1
  },
  "Wecker": {
    "count": 5,
    "lastSeenTime": 1586903561657
  },
  "Ravioli": {
    "count": 7,
    "lastSeenTime": 1586916277356
  },
  "Kolosseum": {
    "count": 9,
    "lastSeenTime": 1586868692820,
    "difficulty": 0.8,
    "difficultyWeight": 15
  },
  "Herde": {
    "count": 5,
    "lastSeenTime": 1586959031017
  },
  "zeigen": {
    "count": 5,
    "lastSeenTime": 1586954989864
  },
  "lächeln": {
    "count": 10,
    "lastSeenTime": 1586839535428,
    "difficulty": 1,
    "difficultyWeight": 9
  },
  "Götterspeise": {
    "count": 7,
    "lastSeenTime": 1586958784537,
    "difficulty": 1,
    "difficultyWeight": 5
  },
  "Injektion": {
    "count": 5,
    "lastSeenTime": 1586956145689
  },
  "Gewürz": {
    "count": 12,
    "lastSeenTime": 1586908982317,
    "difficulty": 1,
    "difficultyWeight": 9
  },
  "Tal": {
    "count": 8,
    "lastSeenTime": 1586913596578
  },
  "Seerose": {
    "count": 8,
    "lastSeenTime": 1586901672587,
    "difficulty": 0.625,
    "difficultyWeight": 8
  },
  "Erdkern": {
    "count": 8,
    "lastSeenTime": 1586959814583,
    "difficulty": 0.8918918918918919,
    "difficultyWeight": 37
  },
  "Motorrad": {
    "count": 8,
    "lastSeenTime": 1586914280738,
    "difficulty": 0.9285714285714286,
    "difficultyWeight": 14
  },
  "Billard": {
    "count": 7,
    "lastSeenTime": 1586915767461
  },
  "Tourist": {
    "count": 5,
    "lastSeenTime": 1586896203705
  },
  "Nachbar": {
    "count": 6,
    "lastSeenTime": 1586890615582
  },
  "Gesicht": {
    "count": 6,
    "lastSeenTime": 1586868886553,
    "difficulty": 0.7058823529411765,
    "difficultyWeight": 17
  },
  "Rick": {
    "count": 10,
    "lastSeenTime": 1586913893763
  },
  "detonieren": {
    "count": 7,
    "lastSeenTime": 1586958100659
  },
  "Elsa": {
    "count": 7,
    "lastSeenTime": 1586916469958,
    "difficulty": 1,
    "difficultyWeight": 9
  },
  "Ehefrau": {
    "count": 2,
    "lastSeenTime": 1586956145689
  },
  "atmen": {
    "count": 5,
    "lastSeenTime": 1586812666102
  },
  "Laser": {
    "count": 5,
    "lastSeenTime": 1586901839909
  },
  "Basketball": {
    "count": 12,
    "lastSeenTime": 1586909325207,
    "difficulty": 1,
    "difficultyWeight": 3
  },
  "Kegelrobbe": {
    "count": 8,
    "lastSeenTime": 1586920346900,
    "difficulty": 1,
    "difficultyWeight": 6
  },
  "Hagel": {
    "count": 6,
    "lastSeenTime": 1586897064138,
    "publicGameCount": 1,
    "difficulty": 0.9473684210526315,
    "difficultyWeight": 19
  },
  "Sitzsack": {
    "count": 9,
    "lastSeenTime": 1586958719765,
    "publicGameCount": 1,
    "difficulty": 0.9024390243902439,
    "difficultyWeight": 41
  },
  "Seuche": {
    "count": 8,
    "lastSeenTime": 1586954664216
  },
  "Bambi": {
    "count": 8,
    "lastSeenTime": 1586900739157
  },
  "Schutzhelm": {
    "count": 6,
    "lastSeenTime": 1586903323590
  },
  "undicht": {
    "count": 7,
    "lastSeenTime": 1586909063921
  },
  "Mund": {
    "count": 6,
    "lastSeenTime": 1586959840559
  },
  "Chips": {
    "count": 5,
    "lastSeenTime": 1586896769426,
    "difficulty": 0.7857142857142857,
    "difficultyWeight": 14
  },
  "Eis": {
    "count": 9,
    "lastSeenTime": 1586897078678,
    "difficulty": 0.7435897435897436,
    "difficultyWeight": 39
  },
  "Kühlschrank": {
    "count": 8,
    "lastSeenTime": 1586959391653,
    "difficulty": 0.8771929824561403,
    "difficultyWeight": 57
  },
  "Ehemann": {
    "count": 10,
    "lastSeenTime": 1586909226093
  },
  "Hammerhai": {
    "count": 11,
    "lastSeenTime": 1586955559002,
    "difficulty": 1,
    "difficultyWeight": 16
  },
  "Lasso": {
    "count": 7,
    "lastSeenTime": 1586902973020,
    "difficulty": 0.8260869565217391,
    "difficultyWeight": 23
  },
  "Elon Musk": {
    "count": 8,
    "lastSeenTime": 1586956273734
  },
  "Steve Jobs": {
    "count": 9,
    "lastSeenTime": 1586908290545
  },
  "Redner": {
    "count": 10,
    "lastSeenTime": 1586920690415,
    "difficulty": 0.896551724137931,
    "difficultyWeight": 29
  },
  "Schaukel": {
    "count": 8,
    "lastSeenTime": 1586913942457
  },
  "Feuerwehrmann": {
    "count": 7,
    "lastSeenTime": 1586891283320,
    "difficulty": 0.9555555555555556,
    "difficultyWeight": 45
  },
  "Atmosphäre": {
    "count": 8,
    "lastSeenTime": 1586904287692,
    "difficulty": 0.9,
    "difficultyWeight": 10
  },
  "Android": {
    "count": 6,
    "lastSeenTime": 1586901554174
  },
  "Frankenstein": {
    "count": 5,
    "lastSeenTime": 1586958748793
  },
  "Neptun": {
    "count": 6,
    "lastSeenTime": 1586954713056
  },
  "Bruch": {
    "count": 5,
    "lastSeenTime": 1586880896100,
    "difficulty": 0.7142857142857143,
    "difficultyWeight": 7
  },
  "Handschlag": {
    "count": 6,
    "lastSeenTime": 1586895467027
  },
  "Wolke": {
    "count": 6,
    "lastSeenTime": 1586954975722
  },
  "Schwarm": {
    "count": 12,
    "lastSeenTime": 1586959633128,
    "publicGameCount": 1,
    "difficulty": 1,
    "difficultyWeight": 9
  },
  "Rettich": {
    "count": 3,
    "lastSeenTime": 1586819449062
  },
  "Kirsche": {
    "count": 5,
    "lastSeenTime": 1586921142411
  },
  "Anime": {
    "count": 8,
    "lastSeenTime": 1586955286110,
    "difficulty": 1,
    "difficultyWeight": 1
  },
  "Wasserschildkröte": {
    "count": 7,
    "lastSeenTime": 1586903673095,
    "difficulty": 0.7777777777777778,
    "difficultyWeight": 18
  },
  "Baguette": {
    "count": 9,
    "lastSeenTime": 1586959242580,
    "difficulty": 0.9,
    "difficultyWeight": 50
  },
  "Badewanne": {
    "count": 6,
    "lastSeenTime": 1586879384676,
    "difficulty": 1,
    "difficultyWeight": 36
  },
  "Maulkorb": {
    "count": 8,
    "lastSeenTime": 1586921043126
  },
  "Bullauge": {
    "count": 8,
    "lastSeenTime": 1586908191616
  },
  "Songtext": {
    "count": 5,
    "lastSeenTime": 1586908418545
  },
  "Fred Feuerstein": {
    "count": 7,
    "lastSeenTime": 1586920843703
  },
  "funkeln": {
    "count": 6,
    "lastSeenTime": 1586869723557,
    "difficulty": 1,
    "difficultyWeight": 5
  },
  "Fossil": {
    "count": 8,
    "lastSeenTime": 1586956453935
  },
  "Widder": {
    "count": 6,
    "lastSeenTime": 1586919477051
  },
  "Bulle": {
    "count": 9,
    "lastSeenTime": 1586920637722
  },
  "Welt": {
    "count": 6,
    "lastSeenTime": 1586914742004,
    "difficulty": 0.5454545454545454,
    "difficultyWeight": 11
  },
  "Surfbrett": {
    "count": 6,
    "lastSeenTime": 1586727413339
  },
  "Benachrichtigung": {
    "count": 6,
    "lastSeenTime": 1586818022119,
    "difficulty": 1,
    "difficultyWeight": 7
  },
  "tief": {
    "count": 6,
    "lastSeenTime": 1586955737202
  },
  "Richter": {
    "count": 6,
    "lastSeenTime": 1586782749678,
    "publicGameCount": 1,
    "difficulty": 0.6666666666666666,
    "difficultyWeight": 3
  },
  "Guillotine": {
    "count": 8,
    "lastSeenTime": 1586915609902,
    "difficulty": 1,
    "difficultyWeight": 7
  },
  "Zahnlücke": {
    "count": 5,
    "lastSeenTime": 1586913571135,
    "difficulty": 0.36363636363636365,
    "difficultyWeight": 11,
    "publicGameCount": 1
  },
  "warm": {
    "count": 6,
    "lastSeenTime": 1586955763543
  },
  "Nachricht": {
    "count": 10,
    "lastSeenTime": 1586914387880,
    "difficulty": 1,
    "difficultyWeight": 10
  },
  "Figur": {
    "count": 4,
    "lastSeenTime": 1586620030988,
    "difficulty": 0.8076923076923077,
    "difficultyWeight": 26
  },
  "Kartoffelpuffer": {
    "count": 6,
    "lastSeenTime": 1586915428040
  },
  "Sprudel": {
    "count": 3,
    "lastSeenTime": 1586828072825
  },
  "Radieschen": {
    "count": 11,
    "lastSeenTime": 1586903975283
  },
  "Pizza": {
    "count": 5,
    "lastSeenTime": 1586807229861,
    "publicGameCount": 1,
    "difficulty": 0.6818181818181818,
    "difficultyWeight": 22
  },
  "Schulter": {
    "count": 7,
    "lastSeenTime": 1586916048294
  },
  "Türkei": {
    "count": 4,
    "lastSeenTime": 1586908304826
  },
  "Kanalisation": {
    "count": 6,
    "lastSeenTime": 1586913988149,
    "difficulty": 1,
    "difficultyWeight": 10
  },
  "Wasserpistole": {
    "count": 6,
    "lastSeenTime": 1586727892454,
    "difficulty": 0.6818181818181818,
    "difficultyWeight": 22
  },
  "Pudding": {
    "count": 6,
    "lastSeenTime": 1586954713056,
    "difficulty": 1,
    "difficultyWeight": 6
  },
  "Benzin": {
    "count": 4,
    "lastSeenTime": 1586706898795,
    "difficulty": 0.9491525423728814,
    "difficultyWeight": 59
  },
  "Spaghetti": {
    "count": 6,
    "lastSeenTime": 1586879335186,
    "difficulty": 0.625,
    "difficultyWeight": 8
  },
  "Micky Maus": {
    "count": 6,
    "lastSeenTime": 1586881158128,
    "difficulty": 1,
    "difficultyWeight": 15,
    "publicGameCount": 1
  },
  "Mark Zuckerberg": {
    "count": 7,
    "lastSeenTime": 1586716859656
  },
  "Elefant": {
    "count": 2,
    "lastSeenTime": 1586714038574,
    "difficulty": 0.875,
    "difficultyWeight": 16
  },
  "Bühne": {
    "count": 3,
    "lastSeenTime": 1586818507852,
    "difficulty": 0.5714285714285714,
    "difficultyWeight": 7
  },
  "Oreo": {
    "count": 5,
    "lastSeenTime": 1586958375030,
    "publicGameCount": 1,
    "difficulty": 0.875,
    "difficultyWeight": 8
  },
  "Nasenbluten": {
    "count": 5,
    "lastSeenTime": 1586827386179,
    "publicGameCount": 1,
    "difficulty": 0.8823529411764706,
    "difficultyWeight": 17
  },
  "Rapper": {
    "count": 6,
    "lastSeenTime": 1586959290196,
    "difficulty": 1,
    "difficultyWeight": 45
  },
  "Bungee Jumping": {
    "count": 7,
    "lastSeenTime": 1586874829202,
    "publicGameCount": 1,
    "difficulty": 0.625,
    "difficultyWeight": 8
  },
  "Abend": {
    "count": 9,
    "lastSeenTime": 1586915115067,
    "difficulty": 0.75,
    "difficultyWeight": 8
  },
  "Crash Bandicoot": {
    "count": 6,
    "lastSeenTime": 1586900136092
  },
  "Welle": {
    "count": 5,
    "lastSeenTime": 1586954727368,
    "difficulty": 0.8695652173913043,
    "difficultyWeight": 23
  },
  "Fass": {
    "count": 7,
    "lastSeenTime": 1586907500591,
    "publicGameCount": 1,
    "difficulty": 0.7142857142857143,
    "difficultyWeight": 7
  },
  "beten": {
    "count": 5,
    "lastSeenTime": 1586859052667,
    "difficulty": 0.8181818181818182,
    "difficultyWeight": 11
  },
  "Marionette": {
    "count": 8,
    "lastSeenTime": 1586956477794
  },
  "Pferdeschwanz": {
    "count": 5,
    "lastSeenTime": 1586904244327,
    "difficulty": 1,
    "difficultyWeight": 8
  },
  "Käfig": {
    "count": 5,
    "lastSeenTime": 1586902702172
  },
  "Dirndl": {
    "count": 7,
    "lastSeenTime": 1586958037635,
    "publicGameCount": 1,
    "difficulty": 0.85,
    "difficultyWeight": 20
  },
  "Fensterbank": {
    "count": 7,
    "lastSeenTime": 1586954834961
  },
  "Laubhaufen": {
    "count": 4,
    "lastSeenTime": 1586959429213,
    "publicGameCount": 1,
    "difficulty": 0.625,
    "difficultyWeight": 8
  },
  "frei schweben": {
    "count": 5,
    "lastSeenTime": 1586915236139,
    "difficulty": 1,
    "difficultyWeight": 10
  },
  "Lineal": {
    "count": 7,
    "lastSeenTime": 1586900054688,
    "difficulty": 1,
    "difficultyWeight": 20,
    "publicGameCount": 1
  },
  "Mr. Meeseeks": {
    "count": 5,
    "lastSeenTime": 1586903133369
  },
  "Volleyball": {
    "count": 9,
    "lastSeenTime": 1586914918478
  },
  "Kanister": {
    "count": 8,
    "lastSeenTime": 1586710393999,
    "difficulty": 0.7142857142857143,
    "difficultyWeight": 7
  },
  "Nasenhaar": {
    "count": 8,
    "lastSeenTime": 1586919917033
  },
  "Apokalypse": {
    "count": 5,
    "lastSeenTime": 1586874424086
  },
  "Hürdenlauf": {
    "count": 4,
    "lastSeenTime": 1586913556716
  },
  "Sprache": {
    "count": 8,
    "lastSeenTime": 1586956501071,
    "difficulty": 1,
    "difficultyWeight": 4
  },
  "Bomberman": {
    "count": 6,
    "lastSeenTime": 1586874456666,
    "difficulty": 0.75,
    "difficultyWeight": 12
  },
  "Labyrinth": {
    "count": 6,
    "lastSeenTime": 1586957848254,
    "difficulty": 1,
    "difficultyWeight": 4
  },
  "Rasierschaum": {
    "count": 5,
    "lastSeenTime": 1586874742120
  },
  "Skrillex": {
    "count": 4,
    "lastSeenTime": 1586914501681
  },
  "Laboratorium": {
    "count": 9,
    "lastSeenTime": 1586958278719
  },
  "Kapitän": {
    "count": 8,
    "lastSeenTime": 1586954900699
  },
  "Chihuahua": {
    "count": 8,
    "lastSeenTime": 1586920761496,
    "publicGameCount": 1
  },
  "giftig": {
    "count": 5,
    "lastSeenTime": 1586814434311
  },
  "winzig": {
    "count": 8,
    "lastSeenTime": 1586919868530
  },
  "Friseur": {
    "count": 7,
    "lastSeenTime": 1586908996510
  },
  "Weizen": {
    "count": 5,
    "lastSeenTime": 1586913670980,
    "difficulty": 0.7,
    "difficultyWeight": 10
  },
  "Trompete": {
    "count": 9,
    "lastSeenTime": 1586958821795,
    "publicGameCount": 1,
    "difficulty": 0.8125,
    "difficultyWeight": 16
  },
  "backen": {
    "count": 8,
    "lastSeenTime": 1586920926700
  },
  "Brennnessel": {
    "count": 9,
    "lastSeenTime": 1586959557107,
    "difficulty": 1,
    "difficultyWeight": 9
  },
  "goldenes Ei": {
    "count": 4,
    "lastSeenTime": 1586880869166,
    "publicGameCount": 1,
    "difficulty": 0.3333333333333333,
    "difficultyWeight": 6
  },
  "Feuerball": {
    "count": 5,
    "lastSeenTime": 1586908079268,
    "difficulty": 0.75,
    "difficultyWeight": 16
  },
  "Maikäfer": {
    "count": 3,
    "lastSeenTime": 1586839864880
  },
  "Clownfisch": {
    "count": 7,
    "lastSeenTime": 1586907361110,
    "difficulty": 0.75,
    "difficultyWeight": 40,
    "publicGameCount": 1
  },
  "Sonnenbrand": {
    "count": 9,
    "lastSeenTime": 1586915960185
  },
  "Vitamin": {
    "count": 5,
    "lastSeenTime": 1586880100518
  },
  "Fastfood": {
    "count": 6,
    "lastSeenTime": 1586958998098
  },
  "Suppe": {
    "count": 9,
    "lastSeenTime": 1586956516194,
    "difficulty": 0.625,
    "difficultyWeight": 8
  },
  "Diener": {
    "count": 4,
    "lastSeenTime": 1586903271433,
    "difficulty": 1,
    "difficultyWeight": 2
  },
  "Diva": {
    "count": 10,
    "lastSeenTime": 1586909335317
  },
  "stark": {
    "count": 9,
    "lastSeenTime": 1586956477794,
    "difficulty": 0.375,
    "difficultyWeight": 8
  },
  "Eiszapfen": {
    "count": 4,
    "lastSeenTime": 1586901471842
  },
  "Physalis": {
    "count": 9,
    "lastSeenTime": 1586920069685
  },
  "Marshmallow": {
    "count": 12,
    "lastSeenTime": 1586959455308
  },
  "Velociraptor": {
    "count": 10,
    "lastSeenTime": 1586915529401,
    "difficulty": 1,
    "difficultyWeight": 3
  },
  "Umweltverschmutzung": {
    "count": 4,
    "lastSeenTime": 1586826994350
  },
  "schüchtern": {
    "count": 7,
    "lastSeenTime": 1586907327679
  },
  "Bayern": {
    "count": 5,
    "lastSeenTime": 1586869155093
  },
  "Gottesanbeterin": {
    "count": 6,
    "lastSeenTime": 1586916027724
  },
  "Pirat": {
    "count": 7,
    "lastSeenTime": 1586920737024
  },
  "Leuchtreklame": {
    "count": 3,
    "lastSeenTime": 1586900446863,
    "difficulty": 1,
    "difficultyWeight": 5
  },
  "Argentinien": {
    "count": 7,
    "lastSeenTime": 1586890826443
  },
  "Zug": {
    "count": 11,
    "lastSeenTime": 1586891319807,
    "publicGameCount": 1,
    "difficulty": 0.9333333333333333,
    "difficultyWeight": 15
  },
  "Ärmel": {
    "count": 5,
    "lastSeenTime": 1586896252833,
    "difficulty": 0.9,
    "difficultyWeight": 10
  },
  "Dschungel": {
    "count": 6,
    "lastSeenTime": 1586832410030,
    "publicGameCount": 1,
    "difficulty": 1,
    "difficultyWeight": 6
  },
  "Zigarette": {
    "count": 12,
    "lastSeenTime": 1586915365103,
    "difficulty": 0.7857142857142857,
    "difficultyWeight": 14
  },
  "Goldfisch": {
    "count": 13,
    "lastSeenTime": 1586907414380,
    "difficulty": 0.7777777777777778,
    "difficultyWeight": 18
  },
  "Junk Food": {
    "count": 6,
    "lastSeenTime": 1586903867648
  },
  "Löwe": {
    "count": 9,
    "lastSeenTime": 1586955726995,
    "difficulty": 1,
    "difficultyWeight": 7
  },
  "Pflaster": {
    "count": 3,
    "lastSeenTime": 1586955249129,
    "difficulty": 1,
    "difficultyWeight": 10
  },
  "zelten": {
    "count": 2,
    "lastSeenTime": 1586958907689,
    "difficulty": 0.8571428571428571,
    "difficultyWeight": 7
  },
  "Gemüse": {
    "count": 7,
    "lastSeenTime": 1586916153776,
    "difficulty": 0.75,
    "difficultyWeight": 12
  },
  "verteidigen": {
    "count": 6,
    "lastSeenTime": 1586921022722
  },
  "Schnauze": {
    "count": 8,
    "lastSeenTime": 1586839274332
  },
  "Dieb": {
    "count": 10,
    "lastSeenTime": 1586832874040
  },
  "Cajon": {
    "count": 5,
    "lastSeenTime": 1586907829080
  },
  "Saxofon": {
    "count": 3,
    "lastSeenTime": 1586955385136
  },
  "Turmalin": {
    "count": 2,
    "lastSeenTime": 1586959314784
  },
  "alt": {
    "count": 6,
    "lastSeenTime": 1586908731467
  },
  "Safari": {
    "count": 5,
    "lastSeenTime": 1586908006992
  },
  "Teich": {
    "count": 8,
    "lastSeenTime": 1586921142411,
    "publicGameCount": 1,
    "difficulty": 0.9166666666666666,
    "difficultyWeight": 12
  },
  "Wildwasserbahn": {
    "count": 8,
    "lastSeenTime": 1586868482944
  },
  "Hula Hoop": {
    "count": 6,
    "lastSeenTime": 1586900001994,
    "difficulty": 0.7692307692307693,
    "difficultyWeight": 13
  },
  "Zahnpasta": {
    "count": 2,
    "lastSeenTime": 1586568563727,
    "publicGameCount": 1,
    "difficulty": 0.5238095238095238,
    "difficultyWeight": 21
  },
  "Rindfleisch": {
    "count": 4,
    "lastSeenTime": 1586860097373
  },
  "Zehnagel": {
    "count": 7,
    "lastSeenTime": 1586914683667
  },
  "Mosaik": {
    "count": 6,
    "lastSeenTime": 1586896117394
  },
  "Rockstar": {
    "count": 4,
    "lastSeenTime": 1586726069070
  },
  "Opfer": {
    "count": 9,
    "lastSeenTime": 1586958697329
  },
  "Kiefer": {
    "count": 9,
    "lastSeenTime": 1586959106188,
    "difficulty": 0.9285714285714286,
    "difficultyWeight": 70
  },
  "Düne": {
    "count": 5,
    "lastSeenTime": 1586915406529,
    "difficulty": 1,
    "difficultyWeight": 17
  },
  "Urlaub": {
    "count": 6,
    "lastSeenTime": 1586901521530,
    "difficulty": 0.5,
    "difficultyWeight": 4
  },
  "Satellit": {
    "count": 5,
    "lastSeenTime": 1586958859906,
    "difficulty": 0,
    "difficultyWeight": 1
  },
  "Bargeld": {
    "count": 3,
    "lastSeenTime": 1586817668923
  },
  "links": {
    "count": 8,
    "lastSeenTime": 1586959840559
  },
  "klettern": {
    "count": 5,
    "lastSeenTime": 1586889515810
  },
  "Graffiti": {
    "count": 8,
    "lastSeenTime": 1586954914968
  },
  "Schrank": {
    "count": 7,
    "lastSeenTime": 1586919941975,
    "difficulty": 0.75,
    "difficultyWeight": 8
  },
  "Apple": {
    "count": 10,
    "lastSeenTime": 1586903370090
  },
  "Flasche": {
    "count": 9,
    "lastSeenTime": 1586958510664
  },
  "Untersetzer": {
    "count": 6,
    "lastSeenTime": 1586901812091,
    "difficulty": 0.88,
    "difficultyWeight": 25
  },
  "Rom": {
    "count": 10,
    "lastSeenTime": 1586908893157
  },
  "Höhlenforscher": {
    "count": 3,
    "lastSeenTime": 1586812916765
  },
  "Picasso": {
    "count": 9,
    "lastSeenTime": 1586897236338,
    "difficulty": 1,
    "difficultyWeight": 6
  },
  "Türsteher": {
    "count": 6,
    "lastSeenTime": 1586916101564
  },
  "Ukulele": {
    "count": 5,
    "lastSeenTime": 1586920797958,
    "difficulty": 0.8571428571428571,
    "difficultyWeight": 14
  },
  "Medaille": {
    "count": 4,
    "lastSeenTime": 1586724733005,
    "difficulty": 0.9830508474576272,
    "difficultyWeight": 59
  },
  "Cousin": {
    "count": 6,
    "lastSeenTime": 1586879212090
  },
  "Bongo": {
    "count": 8,
    "lastSeenTime": 1586881205453,
    "difficulty": 0.45454545454545453,
    "difficultyWeight": 11
  },
  "jung": {
    "count": 10,
    "lastSeenTime": 1586902787660,
    "publicGameCount": 1,
    "difficulty": 0.6666666666666667,
    "difficultyWeight": 9
  },
  "Computer": {
    "count": 7,
    "lastSeenTime": 1586907977622,
    "publicGameCount": 1,
    "difficulty": 0.9285714285714286,
    "difficultyWeight": 14
  },
  "Schotter": {
    "count": 6,
    "lastSeenTime": 1586903429291
  },
  "Flöte": {
    "count": 5,
    "lastSeenTime": 1586908480352,
    "difficulty": 0.9166666666666666,
    "difficultyWeight": 12
  },
  "Eselsohr": {
    "count": 8,
    "lastSeenTime": 1586879959531,
    "publicGameCount": 1,
    "difficulty": 0.9333333333333333,
    "difficultyWeight": 45
  },
  "Rezeptionist": {
    "count": 5,
    "lastSeenTime": 1586874966390
  },
  "Junge": {
    "count": 10,
    "lastSeenTime": 1586958582243
  },
  "Straßenbahn": {
    "count": 6,
    "lastSeenTime": 1586920843703,
    "difficulty": 1,
    "difficultyWeight": 37
  },
  "Murmeln": {
    "count": 7,
    "lastSeenTime": 1586956545635
  },
  "Schere": {
    "count": 9,
    "lastSeenTime": 1586869695200,
    "publicGameCount": 1,
    "difficulty": 0.47058823529411764,
    "difficultyWeight": 17
  },
  "Heckspoiler": {
    "count": 4,
    "lastSeenTime": 1586955958335
  },
  "verdampfen": {
    "count": 10,
    "lastSeenTime": 1586955701781
  },
  "Truthahn": {
    "count": 13,
    "lastSeenTime": 1586958335993
  },
  "Avocado": {
    "count": 9,
    "lastSeenTime": 1586915960185,
    "difficulty": 0.7,
    "difficultyWeight": 10
  },
  "gestresst": {
    "count": 6,
    "lastSeenTime": 1586916023587
  },
  "Sonne": {
    "count": 10,
    "lastSeenTime": 1586958719765,
    "difficulty": 0.5238095238095238,
    "difficultyWeight": 21,
    "publicGameCount": 1
  },
  "Diät": {
    "count": 5,
    "lastSeenTime": 1586896181329
  },
  "traurig": {
    "count": 10,
    "lastSeenTime": 1586955138676,
    "difficulty": 0.6666666666666666,
    "difficultyWeight": 6
  },
  "Elch": {
    "count": 7,
    "lastSeenTime": 1586916027724,
    "difficulty": 0.5,
    "difficultyWeight": 8
  },
  "Luftmatratze": {
    "count": 4,
    "lastSeenTime": 1586840877832,
    "publicGameCount": 1,
    "difficulty": 0.75,
    "difficultyWeight": 8
  },
  "Rubin": {
    "count": 3,
    "lastSeenTime": 1586904161945,
    "difficulty": 0.9285714285714286,
    "difficultyWeight": 14,
    "publicGameCount": 1
  },
  "Waschlappen": {
    "count": 7,
    "lastSeenTime": 1586908504643
  },
  "Eselsbrücke": {
    "count": 4,
    "lastSeenTime": 1586813267167,
    "difficulty": 0.8387096774193549,
    "difficultyWeight": 62
  },
  "Eistee": {
    "count": 9,
    "lastSeenTime": 1586908118744
  },
  "Reise": {
    "count": 10,
    "lastSeenTime": 1586880290728
  },
  "Glücksspiel": {
    "count": 5,
    "lastSeenTime": 1586891062617,
    "difficulty": 1,
    "difficultyWeight": 7
  },
  "Party": {
    "count": 5,
    "lastSeenTime": 1586826796918,
    "difficulty": 0.7674418604651163,
    "difficultyWeight": 43
  },
  "Feuerzeug": {
    "count": 8,
    "lastSeenTime": 1586959207698,
    "difficulty": 1,
    "difficultyWeight": 7
  },
  "Trend": {
    "count": 6,
    "lastSeenTime": 1586895477389
  },
  "Schlüsselbund": {
    "count": 4,
    "lastSeenTime": 1586840879616
  },
  "Zeuge": {
    "count": 2,
    "lastSeenTime": 1586896530203
  },
  "Einschlag": {
    "count": 6,
    "lastSeenTime": 1586955224835
  },
  "Tablet": {
    "count": 12,
    "lastSeenTime": 1586958755749
  },
  "Aussichtspunkt": {
    "count": 12,
    "lastSeenTime": 1586920576642,
    "difficulty": 0.6875,
    "difficultyWeight": 16
  },
  "Käfer": {
    "count": 8,
    "lastSeenTime": 1586890093454
  },
  "Gänseblümchen": {
    "count": 12,
    "lastSeenTime": 1586956263288,
    "difficulty": 1,
    "difficultyWeight": 10
  },
  "ABBA": {
    "count": 5,
    "lastSeenTime": 1586959728501
  },
  "Desoxyribonukleinsäure": {
    "count": 7,
    "lastSeenTime": 1586815682131
  },
  "draußen": {
    "count": 5,
    "lastSeenTime": 1586920119062
  },
  "Präsident": {
    "count": 5,
    "lastSeenTime": 1586815439060
  },
  "Tuba": {
    "count": 6,
    "lastSeenTime": 1586908204031
  },
  "betrügen": {
    "count": 8,
    "lastSeenTime": 1586880201224
  },
  "Kleeblatt": {
    "count": 13,
    "lastSeenTime": 1586959753084,
    "difficulty": 0.7833333333333333,
    "difficultyWeight": 60
  },
  "Fahrrad": {
    "count": 3,
    "lastSeenTime": 1586707253394
  },
  "Seestern": {
    "count": 12,
    "lastSeenTime": 1586921241738,
    "difficulty": 0.5,
    "difficultyWeight": 12,
    "publicGameCount": 1
  },
  "Comic": {
    "count": 5,
    "lastSeenTime": 1586874869836
  },
  "Safe": {
    "count": 6,
    "lastSeenTime": 1586916266718
  },
  "Tschechien": {
    "count": 3,
    "lastSeenTime": 1586707005648,
    "difficulty": 0.7,
    "difficultyWeight": 10
  },
  "Baum": {
    "count": 6,
    "lastSeenTime": 1586908996510,
    "publicGameCount": 1,
    "difficulty": 1,
    "difficultyWeight": 10
  },
  "W-LAN": {
    "count": 7,
    "lastSeenTime": 1586959622590
  },
  "Pokemon": {
    "count": 9,
    "lastSeenTime": 1586957940022
  },
  "Gehirnwäsche": {
    "count": 9,
    "lastSeenTime": 1586904343119
  },
  "Patriot": {
    "count": 12,
    "lastSeenTime": 1586914845805
  },
  "Schwimmbad": {
    "count": 9,
    "lastSeenTime": 1586869648250,
    "publicGameCount": 2,
    "difficulty": 0.75,
    "difficultyWeight": 16
  },
  "Limousine": {
    "count": 6,
    "lastSeenTime": 1586959391653
  },
  "Gesichtsbemalung": {
    "count": 9,
    "lastSeenTime": 1586919941975
  },
  "Paprika": {
    "count": 4,
    "lastSeenTime": 1586895689435
  },
  "Absperrband": {
    "count": 8,
    "lastSeenTime": 1586889393401
  },
  "Kamm": {
    "count": 7,
    "lastSeenTime": 1586904357384,
    "difficulty": 0.85,
    "difficultyWeight": 20
  },
  "Irland": {
    "count": 2,
    "lastSeenTime": 1586889696772
  },
  "Föhn": {
    "count": 7,
    "lastSeenTime": 1586959168567
  },
  "Rose": {
    "count": 7,
    "lastSeenTime": 1586959106188,
    "difficulty": 0.9090909090909091,
    "difficultyWeight": 11
  },
  "Seilrutsche": {
    "count": 10,
    "lastSeenTime": 1586920285874
  },
  "Herbst": {
    "count": 6,
    "lastSeenTime": 1586959533143,
    "difficulty": 1,
    "difficultyWeight": 4
  },
  "Apfel": {
    "count": 9,
    "lastSeenTime": 1586909335317
  },
  "Fahnenstange": {
    "count": 5,
    "lastSeenTime": 1586903221605
  },
  "Skateboard": {
    "count": 6,
    "lastSeenTime": 1586915392140
  },
  "spucken": {
    "count": 3,
    "lastSeenTime": 1586723813838
  },
  "Meerjungfrau": {
    "count": 8,
    "lastSeenTime": 1586914109258,
    "publicGameCount": 1,
    "difficulty": 0.5,
    "difficultyWeight": 8
  },
  "Löwenzahn": {
    "count": 4,
    "lastSeenTime": 1586834155226,
    "difficulty": 0.7777777777777778,
    "difficultyWeight": 9
  },
  "Armbanduhr": {
    "count": 9,
    "lastSeenTime": 1586909311071,
    "publicGameCount": 2,
    "difficulty": 0.6666666666666666,
    "difficultyWeight": 6
  },
  "Fabrik": {
    "count": 2,
    "lastSeenTime": 1586633994738
  },
  "Berühmtheit": {
    "count": 6,
    "lastSeenTime": 1586955938117
  },
  "NASCAR": {
    "count": 4,
    "lastSeenTime": 1586713655448
  },
  "Urknall": {
    "count": 4,
    "lastSeenTime": 1586812433787
  },
  "Nest": {
    "count": 7,
    "lastSeenTime": 1586919750850,
    "difficulty": 0.8571428571428571,
    "difficultyWeight": 14
  },
  "Bernstein": {
    "count": 13,
    "lastSeenTime": 1586959585641,
    "difficulty": 1,
    "difficultyWeight": 4
  },
  "Auge": {
    "count": 19,
    "lastSeenTime": 1586920265574,
    "difficulty": 0.5789473684210527,
    "difficultyWeight": 19,
    "publicGameCount": 1
  },
  "Skispringen": {
    "count": 4,
    "lastSeenTime": 1586954588805
  },
  "Sphinx": {
    "count": 4,
    "lastSeenTime": 1586881119264
  },
  "Wind": {
    "count": 7,
    "lastSeenTime": 1586919968806,
    "difficulty": 1,
    "difficultyWeight": 16
  },
  "Piratenschiff": {
    "count": 4,
    "lastSeenTime": 1586900828046,
    "difficulty": 0.9459459459459459,
    "difficultyWeight": 37
  },
  "Twitter": {
    "count": 10,
    "lastSeenTime": 1586958848321,
    "publicGameCount": 1,
    "difficulty": 0.8,
    "difficultyWeight": 10
  },
  "genau": {
    "count": 8,
    "lastSeenTime": 1586826669362,
    "difficulty": 1,
    "difficultyWeight": 3
  },
  "Pelikan": {
    "count": 5,
    "lastSeenTime": 1586954698858,
    "difficulty": 1,
    "difficultyWeight": 2
  },
  "Katastrophe": {
    "count": 4,
    "lastSeenTime": 1586914754340
  },
  "Hotdog": {
    "count": 5,
    "lastSeenTime": 1586908838335
  },
  "Sport": {
    "count": 6,
    "lastSeenTime": 1586813898884
  },
  "Toilette": {
    "count": 5,
    "lastSeenTime": 1586873774235,
    "difficulty": 0.9166666666666666,
    "difficultyWeight": 12
  },
  "Schranke": {
    "count": 6,
    "lastSeenTime": 1586903059681
  },
  "lustig": {
    "count": 6,
    "lastSeenTime": 1586858614527
  },
  "Fleischbällchen": {
    "count": 7,
    "lastSeenTime": 1586958385310,
    "difficulty": 1,
    "difficultyWeight": 6
  },
  "Strudel": {
    "count": 5,
    "lastSeenTime": 1586959193290
  },
  "Spiegelei": {
    "count": 7,
    "lastSeenTime": 1586869103569,
    "difficulty": 0.6666666666666667,
    "difficultyWeight": 9
  },
  "Einhorn": {
    "count": 10,
    "lastSeenTime": 1586901028464,
    "publicGameCount": 1,
    "difficulty": 0.75,
    "difficultyWeight": 8
  },
  "Rücken": {
    "count": 6,
    "lastSeenTime": 1586959417006,
    "difficulty": 1,
    "difficultyWeight": 36
  },
  "pflügen": {
    "count": 5,
    "lastSeenTime": 1586908054437
  },
  "Goofy": {
    "count": 9,
    "lastSeenTime": 1586959789851
  },
  "Seifenoper": {
    "count": 4,
    "lastSeenTime": 1586902616627
  },
  "Kopflaus": {
    "count": 3,
    "lastSeenTime": 1586879608242
  },
  "Traumfänger": {
    "count": 5,
    "lastSeenTime": 1586903097380
  },
  "Kernspintomografie": {
    "count": 6,
    "lastSeenTime": 1586921270600
  },
  "klingen": {
    "count": 5,
    "lastSeenTime": 1586914243836
  },
  "Kaulquappe": {
    "count": 6,
    "lastSeenTime": 1586913698511,
    "difficulty": 0.8823529411764706,
    "difficultyWeight": 17
  },
  "vorspulen": {
    "count": 6,
    "lastSeenTime": 1586913816802
  },
  "Augenlid": {
    "count": 10,
    "lastSeenTime": 1586957950175
  },
  "Panda": {
    "count": 11,
    "lastSeenTime": 1586920637722
  },
  "Singapur": {
    "count": 5,
    "lastSeenTime": 1586840929412
  },
  "Zauberer": {
    "count": 6,
    "lastSeenTime": 1586955238992
  },
  "Himbeere": {
    "count": 10,
    "lastSeenTime": 1586909016736
  },
  "Cerberus": {
    "count": 7,
    "lastSeenTime": 1586959830012
  },
  "Pfirsich": {
    "count": 5,
    "lastSeenTime": 1586806870118
  },
  "Dachs": {
    "count": 2,
    "lastSeenTime": 1586914418213
  },
  "bleichen": {
    "count": 6,
    "lastSeenTime": 1586881168379
  },
  "Photoshop": {
    "count": 7,
    "lastSeenTime": 1586899941163
  },
  "Hut": {
    "count": 7,
    "lastSeenTime": 1586955344691
  },
  "Text": {
    "count": 6,
    "lastSeenTime": 1586959269543
  },
  "Jesus": {
    "count": 3,
    "lastSeenTime": 1586880616380,
    "difficulty": 0.8888888888888888,
    "difficultyWeight": 9
  },
  "stehen": {
    "count": 8,
    "lastSeenTime": 1586957974879
  },
  "Chinatown": {
    "count": 4,
    "lastSeenTime": 1586832383649
  },
  "Tausendfüßer": {
    "count": 11,
    "lastSeenTime": 1586903624140,
    "difficulty": 0.6,
    "difficultyWeight": 10
  },
  "Schlitten": {
    "count": 8,
    "lastSeenTime": 1586873964321,
    "difficulty": 0.7647058823529411,
    "difficultyWeight": 34
  },
  "Rakete": {
    "count": 12,
    "lastSeenTime": 1586955838030,
    "difficulty": 1,
    "difficultyWeight": 104
  },
  "Glatze": {
    "count": 4,
    "lastSeenTime": 1586891242624,
    "difficulty": 0.7441860465116279,
    "difficultyWeight": 43
  },
  "Kopie": {
    "count": 8,
    "lastSeenTime": 1586955392320
  },
  "älter": {
    "count": 10,
    "lastSeenTime": 1586955634694
  },
  "Weihnachtsbaum": {
    "count": 7,
    "lastSeenTime": 1586913606771,
    "difficulty": 0.6,
    "difficultyWeight": 5
  },
  "Ruder": {
    "count": 5,
    "lastSeenTime": 1586897420912,
    "difficulty": 0.8333333333333334,
    "difficultyWeight": 24
  },
  "Ecke": {
    "count": 7,
    "lastSeenTime": 1586920576642
  },
  "Bohne": {
    "count": 8,
    "lastSeenTime": 1586958998098
  },
  "Apotheker": {
    "count": 9,
    "lastSeenTime": 1586915566245,
    "difficulty": 0.8333333333333334,
    "difficultyWeight": 18
  },
  "Scherzkeks": {
    "count": 5,
    "lastSeenTime": 1586955654939
  },
  "Küste": {
    "count": 9,
    "lastSeenTime": 1586900749848,
    "publicGameCount": 1,
    "difficulty": 0.625,
    "difficultyWeight": 8
  },
  "anzünden": {
    "count": 8,
    "lastSeenTime": 1586958027091,
    "difficulty": 0.8,
    "difficultyWeight": 20
  },
  "googeln": {
    "count": 7,
    "lastSeenTime": 1586920133284,
    "difficulty": 1,
    "difficultyWeight": 6
  },
  "wackeln": {
    "count": 4,
    "lastSeenTime": 1586873308659
  },
  "Roter Teppich": {
    "count": 3,
    "lastSeenTime": 1586890539988,
    "difficulty": 1,
    "difficultyWeight": 3
  },
  "Scheune": {
    "count": 6,
    "lastSeenTime": 1586920034941
  },
  "Schlachter": {
    "count": 5,
    "lastSeenTime": 1586873642917
  },
  "Vogelscheuche": {
    "count": 7,
    "lastSeenTime": 1586959417006,
    "publicGameCount": 1,
    "difficulty": 0.5714285714285714,
    "difficultyWeight": 14
  },
  "Sauerstoff": {
    "count": 5,
    "lastSeenTime": 1586897343815
  },
  "Ente": {
    "count": 4,
    "lastSeenTime": 1586914831496,
    "difficulty": 0.9333333333333333,
    "difficultyWeight": 15
  },
  "blind": {
    "count": 6,
    "lastSeenTime": 1586859567294,
    "difficulty": 0.7222222222222222,
    "difficultyWeight": 18
  },
  "Blasebalg": {
    "count": 9,
    "lastSeenTime": 1586920483192,
    "difficulty": 1,
    "difficultyWeight": 6
  },
  "Kathedrale": {
    "count": 7,
    "lastSeenTime": 1586915624112
  },
  "Altglascontainer": {
    "count": 2,
    "lastSeenTime": 1586619769643
  },
  "Poster": {
    "count": 11,
    "lastSeenTime": 1586955913834,
    "difficulty": 0.5,
    "difficultyWeight": 8
  },
  "rot": {
    "count": 6,
    "lastSeenTime": 1586900776963
  },
  "Windel": {
    "count": 5,
    "lastSeenTime": 1586726463833
  },
  "Bösewicht": {
    "count": 4,
    "lastSeenTime": 1586903915052
  },
  "Bettwanze": {
    "count": 7,
    "lastSeenTime": 1586915006438
  },
  "schwingen": {
    "count": 10,
    "lastSeenTime": 1586956009184,
    "difficulty": 0.9166666666666666,
    "difficultyWeight": 12
  },
  "Kaffee": {
    "count": 9,
    "lastSeenTime": 1586908494479
  },
  "Tennis": {
    "count": 5,
    "lastSeenTime": 1586710847152
  },
  "Priester": {
    "count": 7,
    "lastSeenTime": 1586955517664
  },
  "elektrisch": {
    "count": 3,
    "lastSeenTime": 1586914280738
  },
  "Schneeflocke": {
    "count": 10,
    "lastSeenTime": 1586921032863,
    "difficulty": 0.5714285714285714,
    "difficultyWeight": 7
  },
  "Bus": {
    "count": 7,
    "lastSeenTime": 1586956453935
  },
  "Kindergarten": {
    "count": 5,
    "lastSeenTime": 1586900125985
  },
  "Goblin": {
    "count": 4,
    "lastSeenTime": 1586869879731
  },
  "Revolver": {
    "count": 9,
    "lastSeenTime": 1586956489403
  },
  "Versteck": {
    "count": 6,
    "lastSeenTime": 1586902853216
  },
  "Polarlicht": {
    "count": 12,
    "lastSeenTime": 1586907486194,
    "difficulty": 0.8431372549019608,
    "difficultyWeight": 51
  },
  "Speer": {
    "count": 8,
    "lastSeenTime": 1586913581904
  },
  "befehlen": {
    "count": 5,
    "lastSeenTime": 1586726791412
  },
  "Beute": {
    "count": 5,
    "lastSeenTime": 1586879469524
  },
  "Möhre": {
    "count": 10,
    "lastSeenTime": 1586954713056
  },
  "MTV": {
    "count": 13,
    "lastSeenTime": 1586914373601
  },
  "Nutella": {
    "count": 4,
    "lastSeenTime": 1586869069583,
    "difficulty": 1,
    "difficultyWeight": 2
  },
  "Grinsen": {
    "count": 6,
    "lastSeenTime": 1586959121500,
    "difficulty": 1,
    "difficultyWeight": 8
  },
  "saufen": {
    "count": 3,
    "lastSeenTime": 1586726288386
  },
  "Koralle": {
    "count": 14,
    "lastSeenTime": 1586955224835
  },
  "Ende": {
    "count": 3,
    "lastSeenTime": 1586813572686,
    "difficulty": 0.6666666666666666,
    "difficultyWeight": 3
  },
  "Halbinsel": {
    "count": 6,
    "lastSeenTime": 1586869576824
  },
  "Wasserfall": {
    "count": 5,
    "lastSeenTime": 1586916565022,
    "difficulty": 0.6666666666666666,
    "difficultyWeight": 9,
    "publicGameCount": 1
  },
  "Fuß": {
    "count": 3,
    "lastSeenTime": 1586909104982
  },
  "Thermometer": {
    "count": 4,
    "lastSeenTime": 1586897053852
  },
  "Chat": {
    "count": 6,
    "lastSeenTime": 1586838997852,
    "difficulty": 0.7058823529411765,
    "difficultyWeight": 17
  },
  "Werbespot": {
    "count": 4,
    "lastSeenTime": 1586726387745
  },
  "Gebrauchtwagenhändler": {
    "count": 9,
    "lastSeenTime": 1586920538046,
    "difficulty": 1,
    "difficultyWeight": 10
  },
  "Fisch": {
    "count": 11,
    "lastSeenTime": 1586916404987,
    "difficulty": 1,
    "difficultyWeight": 36
  },
  "zerreißen": {
    "count": 5,
    "lastSeenTime": 1586919697403,
    "difficulty": 1,
    "difficultyWeight": 10
  },
  "Schaufensterpuppe": {
    "count": 7,
    "lastSeenTime": 1586916382544
  },
  "Tapete": {
    "count": 11,
    "lastSeenTime": 1586903246007,
    "difficulty": 0.6111111111111112,
    "difficultyWeight": 18
  },
  "Ober": {
    "count": 8,
    "lastSeenTime": 1586955809521,
    "difficulty": 1,
    "difficultyWeight": 4
  },
  "Diplom": {
    "count": 7,
    "lastSeenTime": 1586916252219
  },
  "Strahlung": {
    "count": 8,
    "lastSeenTime": 1586900895312
  },
  "stechen": {
    "count": 10,
    "lastSeenTime": 1586958642478
  },
  "Espresso": {
    "count": 5,
    "lastSeenTime": 1586909006610,
    "difficulty": 0.8333333333333334,
    "difficultyWeight": 18
  },
  "Rutsche": {
    "count": 6,
    "lastSeenTime": 1586868015613
  },
  "skribbl.io": {
    "count": 4,
    "lastSeenTime": 1586832435328
  },
  "kurz": {
    "count": 10,
    "lastSeenTime": 1586907668803
  },
  "Wörterbuch": {
    "count": 4,
    "lastSeenTime": 1586727362775,
    "difficulty": 1,
    "difficultyWeight": 17
  },
  "Mont Blanc": {
    "count": 6,
    "lastSeenTime": 1586815159300
  },
  "Hamster": {
    "count": 7,
    "lastSeenTime": 1586869613566
  },
  "Artist": {
    "count": 3,
    "lastSeenTime": 1586914869480
  },
  "Champion": {
    "count": 3,
    "lastSeenTime": 1586817823665,
    "difficulty": 0.875,
    "difficultyWeight": 8
  },
  "Addition": {
    "count": 6,
    "lastSeenTime": 1586956253026
  },
  "Haselnuss": {
    "count": 6,
    "lastSeenTime": 1586921096637
  },
  "Oktoberfest": {
    "count": 8,
    "lastSeenTime": 1586955862348
  },
  "Kürbislaterne": {
    "count": 6,
    "lastSeenTime": 1586900927708
  },
  "Glücksrad": {
    "count": 9,
    "lastSeenTime": 1586818937722
  },
  "schielen": {
    "count": 3,
    "lastSeenTime": 1586596228156
  },
  "zerren": {
    "count": 10,
    "lastSeenTime": 1586955409372
  },
  "Hockey": {
    "count": 4,
    "lastSeenTime": 1586900226190,
    "difficulty": 0.8636363636363636,
    "difficultyWeight": 22
  },
  "Keller": {
    "count": 3,
    "lastSeenTime": 1586839939626
  },
  "Playstation": {
    "count": 9,
    "lastSeenTime": 1586915493381,
    "difficulty": 0.7692307692307693,
    "difficultyWeight": 13
  },
  "Bugs Bunny": {
    "count": 5,
    "lastSeenTime": 1586954698858
  },
  "Spender": {
    "count": 7,
    "lastSeenTime": 1586784575401
  },
  "Flaschendrehen": {
    "count": 6,
    "lastSeenTime": 1586879929914,
    "publicGameCount": 1,
    "difficulty": 0.7,
    "difficultyWeight": 10
  },
  "Stuhlbein": {
    "count": 7,
    "lastSeenTime": 1586900412247,
    "difficulty": 1,
    "difficultyWeight": 32
  },
  "angreifen": {
    "count": 4,
    "lastSeenTime": 1586827564451
  },
  "Popeye": {
    "count": 6,
    "lastSeenTime": 1586873629254
  },
  "Verbrecher": {
    "count": 8,
    "lastSeenTime": 1586916342434,
    "difficulty": 0.8,
    "difficultyWeight": 5
  },
  "Federmäppchen": {
    "count": 5,
    "lastSeenTime": 1586902813258
  },
  "BMX": {
    "count": 6,
    "lastSeenTime": 1586916266718
  },
  "Chamäleon": {
    "count": 8,
    "lastSeenTime": 1586880818922,
    "difficulty": 0.8571428571428571,
    "difficultyWeight": 14
  },
  "Europa": {
    "count": 5,
    "lastSeenTime": 1586873585720
  },
  "Sattelschlepper": {
    "count": 6,
    "lastSeenTime": 1586828072825
  },
  "Magma": {
    "count": 5,
    "lastSeenTime": 1586717707336,
    "difficulty": 1,
    "difficultyWeight": 8
  },
  "Beule": {
    "count": 6,
    "lastSeenTime": 1586859604461
  },
  "Jalapeno": {
    "count": 5,
    "lastSeenTime": 1586869197437
  },
  "Detektiv": {
    "count": 7,
    "lastSeenTime": 1586913497512
  },
  "Ei": {
    "count": 8,
    "lastSeenTime": 1586859308294,
    "difficulty": 0.6666666666666666,
    "difficultyWeight": 48
  },
  "Storch": {
    "count": 3,
    "lastSeenTime": 1586895909639
  },
  "Gasse": {
    "count": 4,
    "lastSeenTime": 1586958873120,
    "difficulty": 1,
    "difficultyWeight": 9
  },
  "Champagner": {
    "count": 7,
    "lastSeenTime": 1586920676259
  },
  "Werbung": {
    "count": 7,
    "lastSeenTime": 1586914094806,
    "difficulty": 0.9565217391304348,
    "difficultyWeight": 46
  },
  "Hulk": {
    "count": 3,
    "lastSeenTime": 1586880213361,
    "difficulty": 1,
    "difficultyWeight": 8
  },
  "Mafia": {
    "count": 7,
    "lastSeenTime": 1586920153587,
    "difficulty": 1,
    "difficultyWeight": 8
  },
  "Strandkorb": {
    "count": 9,
    "lastSeenTime": 1586956016971
  },
  "kitzeln": {
    "count": 8,
    "lastSeenTime": 1586920883005
  },
  "Mexiko": {
    "count": 6,
    "lastSeenTime": 1586879731749
  },
  "Hochzeitskutsche": {
    "count": 5,
    "lastSeenTime": 1586957989157,
    "difficulty": 1,
    "difficultyWeight": 9
  },
  "Pastete": {
    "count": 8,
    "lastSeenTime": 1586819518500
  },
  "Schilf": {
    "count": 9,
    "lastSeenTime": 1586897164255
  },
  "Intel": {
    "count": 11,
    "lastSeenTime": 1586919413397
  },
  "Weihnachten": {
    "count": 3,
    "lastSeenTime": 1586880506505,
    "difficulty": 1,
    "difficultyWeight": 12
  },
  "Dose": {
    "count": 7,
    "lastSeenTime": 1586920237009
  },
  "Luchs": {
    "count": 3,
    "lastSeenTime": 1586595496641,
    "difficulty": 1,
    "difficultyWeight": 11
  },
  "Vogelstrauß": {
    "count": 7,
    "lastSeenTime": 1586902713347,
    "difficulty": 0.9090909090909091,
    "difficultyWeight": 11
  },
  "Stirn": {
    "count": 2,
    "lastSeenTime": 1586890418363
  },
  "Bar": {
    "count": 9,
    "lastSeenTime": 1586903925357
  },
  "Lama": {
    "count": 9,
    "lastSeenTime": 1586956501071
  },
  "Orchidee": {
    "count": 3,
    "lastSeenTime": 1586899991881
  },
  "berühmt": {
    "count": 9,
    "lastSeenTime": 1586909119149
  },
  "Boris Becker": {
    "count": 2,
    "lastSeenTime": 1586580996457,
    "difficulty": 1,
    "difficultyWeight": 5
  },
  "Katana": {
    "count": 4,
    "lastSeenTime": 1586724355196
  },
  "Riesenrad": {
    "count": 5,
    "lastSeenTime": 1586955681359
  },
  "Social Media": {
    "count": 5,
    "lastSeenTime": 1586896605134
  },
  "Voodoo": {
    "count": 10,
    "lastSeenTime": 1586901206044,
    "difficulty": 1,
    "difficultyWeight": 6
  },
  "Radar": {
    "count": 5,
    "lastSeenTime": 1586900301076,
    "difficulty": 1,
    "difficultyWeight": 5
  },
  "Tortenheber": {
    "count": 10,
    "lastSeenTime": 1586913827190
  },
  "Bauernhof": {
    "count": 3,
    "lastSeenTime": 1586901801777
  },
  "Segway": {
    "count": 3,
    "lastSeenTime": 1586920900377,
    "difficulty": 0.7575757575757576,
    "difficultyWeight": 33
  },
  "Ringelblume": {
    "count": 4,
    "lastSeenTime": 1586914903455
  },
  "Trophäe": {
    "count": 7,
    "lastSeenTime": 1586900881142
  },
  "Atem": {
    "count": 5,
    "lastSeenTime": 1586958012186
  },
  "Oase": {
    "count": 7,
    "lastSeenTime": 1586896127512
  },
  "Weinglas": {
    "count": 3,
    "lastSeenTime": 1586895726867,
    "difficulty": 1,
    "difficultyWeight": 3
  },
  "Reinkarnation": {
    "count": 14,
    "lastSeenTime": 1586916637868
  },
  "Matrjoschka": {
    "count": 6,
    "lastSeenTime": 1586915664822
  },
  "Regenmantel": {
    "count": 6,
    "lastSeenTime": 1586955737202
  },
  "Peppa Pig": {
    "count": 6,
    "lastSeenTime": 1586818357030
  },
  "Tastatur": {
    "count": 9,
    "lastSeenTime": 1586919394915,
    "difficulty": 1,
    "difficultyWeight": 8
  },
  "Samstag": {
    "count": 13,
    "lastSeenTime": 1586955370991,
    "difficulty": 1,
    "difficultyWeight": 8
  },
  "Killerwal": {
    "count": 6,
    "lastSeenTime": 1586895562695
  },
  "Sternfrucht": {
    "count": 9,
    "lastSeenTime": 1586921256075,
    "difficulty": 0.6666666666666666,
    "difficultyWeight": 6
  },
  "Muskel": {
    "count": 8,
    "lastSeenTime": 1586959695898
  },
  "erschreckend": {
    "count": 4,
    "lastSeenTime": 1586709223746,
    "publicGameCount": 1,
    "difficulty": 0.8333333333333334,
    "difficultyWeight": 6
  },
  "Allergie": {
    "count": 6,
    "lastSeenTime": 1586956352310,
    "difficulty": 0.42857142857142855,
    "difficultyWeight": 7
  },
  "Himmel": {
    "count": 10,
    "lastSeenTime": 1586919736561,
    "difficulty": 0.94,
    "difficultyWeight": 50
  },
  "Oberteil": {
    "count": 7,
    "lastSeenTime": 1586958642478
  },
  "Kreide": {
    "count": 9,
    "lastSeenTime": 1586958973713
  },
  "Kermit": {
    "count": 6,
    "lastSeenTime": 1586869916885
  },
  "Eis am Stiel": {
    "count": 4,
    "lastSeenTime": 1586914134487
  },
  "Staffelei": {
    "count": 8,
    "lastSeenTime": 1586909063921,
    "publicGameCount": 1,
    "difficulty": 1,
    "difficultyWeight": 32
  },
  "Kissenschlacht": {
    "count": 7,
    "lastSeenTime": 1586914590565
  },
  "Falle": {
    "count": 6,
    "lastSeenTime": 1586900828046,
    "difficulty": 0.8421052631578947,
    "difficultyWeight": 38
  },
  "Korb": {
    "count": 10,
    "lastSeenTime": 1586920346900
  },
  "Limette": {
    "count": 6,
    "lastSeenTime": 1586954664216,
    "difficulty": 0.9310344827586207,
    "difficultyWeight": 29
  },
  "Rauch": {
    "count": 8,
    "lastSeenTime": 1586914820916
  },
  "Helium": {
    "count": 6,
    "lastSeenTime": 1586916550509
  },
  "Ahorn": {
    "count": 5,
    "lastSeenTime": 1586890315125
  },
  "Klobrille": {
    "count": 7,
    "lastSeenTime": 1586903271433,
    "publicGameCount": 1,
    "difficulty": 0.875,
    "difficultyWeight": 16
  },
  "falten": {
    "count": 5,
    "lastSeenTime": 1586868317304,
    "difficulty": 0.7142857142857143,
    "difficultyWeight": 14
  },
  "Preisschild": {
    "count": 4,
    "lastSeenTime": 1586908189842
  },
  "Shampoo": {
    "count": 5,
    "lastSeenTime": 1586596511045
  },
  "Hängematte": {
    "count": 4,
    "lastSeenTime": 1586913783575,
    "difficulty": 0.4444444444444444,
    "difficultyWeight": 9
  },
  "Monaco": {
    "count": 9,
    "lastSeenTime": 1586955385136
  },
  "Laptop": {
    "count": 11,
    "lastSeenTime": 1586956545635,
    "difficulty": 0.42857142857142855,
    "difficultyWeight": 7
  },
  "Banane": {
    "count": 6,
    "lastSeenTime": 1586873408581
  },
  "Pac-Man": {
    "count": 7,
    "lastSeenTime": 1586889858351,
    "difficulty": 0.9375,
    "difficultyWeight": 16
  },
  "Radio": {
    "count": 3,
    "lastSeenTime": 1586727102331,
    "difficulty": 0.7777777777777778,
    "difficultyWeight": 36
  },
  "Mörder": {
    "count": 3,
    "lastSeenTime": 1586727017449,
    "difficulty": 1,
    "difficultyWeight": 6
  },
  "Sanduhr": {
    "count": 7,
    "lastSeenTime": 1586826882193,
    "difficulty": 0.7,
    "difficultyWeight": 10
  },
  "Jay Z": {
    "count": 4,
    "lastSeenTime": 1586879531452
  },
  "Einkaufswagen": {
    "count": 8,
    "lastSeenTime": 1586816175412,
    "difficulty": 1,
    "difficultyWeight": 8
  },
  "Bluterguss": {
    "count": 7,
    "lastSeenTime": 1586956324955,
    "publicGameCount": 1,
    "difficulty": 1,
    "difficultyWeight": 7
  },
  "Green Lantern": {
    "count": 7,
    "lastSeenTime": 1586959595986
  },
  "Quartal": {
    "count": 2,
    "lastSeenTime": 1586716898363
  },
  "Spinne": {
    "count": 8,
    "lastSeenTime": 1586913633876,
    "difficulty": 0.5263157894736842,
    "difficultyWeight": 19
  },
  "Pfannkuchen": {
    "count": 9,
    "lastSeenTime": 1586956545635,
    "difficulty": 0.9607843137254902,
    "difficultyWeight": 51
  },
  "Rippe": {
    "count": 14,
    "lastSeenTime": 1586958471219,
    "difficulty": 0.9166666666666666,
    "difficultyWeight": 12
  },
  "schwarz": {
    "count": 5,
    "lastSeenTime": 1586954999969,
    "publicGameCount": 1,
    "difficulty": 0.6666666666666666,
    "difficultyWeight": 3
  },
  "Silberbesteck": {
    "count": 8,
    "lastSeenTime": 1586816892701
  },
  "Russland": {
    "count": 2,
    "lastSeenTime": 1586568214189,
    "difficulty": 1,
    "difficultyWeight": 8
  },
  "Sumpf": {
    "count": 7,
    "lastSeenTime": 1586915428040
  },
  "Vogelbad": {
    "count": 6,
    "lastSeenTime": 1586903781874
  },
  "Torte": {
    "count": 6,
    "lastSeenTime": 1586914727219,
    "difficulty": 0.8,
    "difficultyWeight": 20
  },
  "Klaue": {
    "count": 4,
    "lastSeenTime": 1586919394915
  },
  "Sherlock Holmes": {
    "count": 7,
    "lastSeenTime": 1586891268927,
    "difficulty": 1,
    "difficultyWeight": 8
  },
  "Kronleuchter": {
    "count": 4,
    "lastSeenTime": 1586900076992,
    "difficulty": 0.5714285714285714,
    "difficultyWeight": 7
  },
  "Muskelkater": {
    "count": 9,
    "lastSeenTime": 1586958683104,
    "publicGameCount": 1,
    "difficulty": 0.875,
    "difficultyWeight": 8
  },
  "Land": {
    "count": 5,
    "lastSeenTime": 1586896155897
  },
  "Käse": {
    "count": 2,
    "lastSeenTime": 1586712115480,
    "difficulty": 0.75,
    "difficultyWeight": 4
  },
  "Vatikan": {
    "count": 7,
    "lastSeenTime": 1586880423720,
    "publicGameCount": 1
  },
  "Südpol": {
    "count": 7,
    "lastSeenTime": 1586916317911,
    "difficulty": 1,
    "difficultyWeight": 12
  },
  "feiern": {
    "count": 12,
    "lastSeenTime": 1586955286110
  },
  "Netzwerk": {
    "count": 8,
    "lastSeenTime": 1586908672662,
    "difficulty": 0.7272727272727273,
    "difficultyWeight": 11
  },
  "Taschenlampe": {
    "count": 5,
    "lastSeenTime": 1586880869166,
    "difficulty": 0.5238095238095238,
    "difficultyWeight": 21
  },
  "Spieß": {
    "count": 10,
    "lastSeenTime": 1586919384885,
    "difficulty": 0.2222222222222222,
    "difficultyWeight": 9
  },
  "Wohnzimmer": {
    "count": 3,
    "lastSeenTime": 1586569967340,
    "publicGameCount": 1
  },
  "königlich": {
    "count": 8,
    "lastSeenTime": 1586921180981
  },
  "Strand": {
    "count": 8,
    "lastSeenTime": 1586954925078,
    "difficulty": 0.8571428571428571,
    "difficultyWeight": 28
  },
  "Bärenfalle": {
    "count": 8,
    "lastSeenTime": 1586901239971
  },
  "Abfall": {
    "count": 8,
    "lastSeenTime": 1586807202609,
    "difficulty": 0.84,
    "difficultyWeight": 25,
    "publicGameCount": 1
  },
  "LKW-Fahrer": {
    "count": 6,
    "lastSeenTime": 1586915354278,
    "difficulty": 0.875,
    "difficultyWeight": 72
  },
  "Wand": {
    "count": 6,
    "lastSeenTime": 1586908175585
  },
  "schwimmen": {
    "count": 5,
    "lastSeenTime": 1586916127462,
    "difficulty": 0.9722222222222222,
    "difficultyWeight": 36
  },
  "Peitsche": {
    "count": 5,
    "lastSeenTime": 1586914893907,
    "difficulty": 1,
    "difficultyWeight": 52
  },
  "Leichtsinn": {
    "count": 8,
    "lastSeenTime": 1586903925357
  },
  "Luftkissenboot": {
    "count": 7,
    "lastSeenTime": 1586958485799
  },
  "Einhornwal": {
    "count": 8,
    "lastSeenTime": 1586919423610
  },
  "Iron Man": {
    "count": 5,
    "lastSeenTime": 1586890464755,
    "publicGameCount": 1
  },
  "Schlüssel": {
    "count": 5,
    "lastSeenTime": 1586834021563,
    "difficulty": 0.4444444444444444,
    "difficultyWeight": 9
  },
  "Basis": {
    "count": 12,
    "lastSeenTime": 1586874906969
  },
  "Pickel": {
    "count": 8,
    "lastSeenTime": 1586959695898
  },
  "unendlich": {
    "count": 7,
    "lastSeenTime": 1586955691477,
    "difficulty": 0.4166666666666667,
    "difficultyWeight": 12
  },
  "Gru": {
    "count": 9,
    "lastSeenTime": 1586919797630
  },
  "Trittstufe": {
    "count": 6,
    "lastSeenTime": 1586957855578
  },
  "Handfläche": {
    "count": 8,
    "lastSeenTime": 1586919640209
  },
  "jonglieren": {
    "count": 8,
    "lastSeenTime": 1586958658325
  },
  "Korallenriff": {
    "count": 5,
    "lastSeenTime": 1586916127462,
    "publicGameCount": 1,
    "difficulty": 0.75,
    "difficultyWeight": 8
  },
  "Achselhöhle": {
    "count": 7,
    "lastSeenTime": 1586916680436
  },
  "Taco": {
    "count": 11,
    "lastSeenTime": 1586920858569,
    "difficulty": 0.7560975609756098,
    "difficultyWeight": 41,
    "publicGameCount": 1
  },
  "Violine": {
    "count": 6,
    "lastSeenTime": 1586806873561,
    "difficulty": 0.38461538461538464,
    "difficultyWeight": 13
  },
  "Bahnschiene": {
    "count": 11,
    "lastSeenTime": 1586889379582
  },
  "mürrisch": {
    "count": 6,
    "lastSeenTime": 1586919797630,
    "difficulty": 0.5625,
    "difficultyWeight": 16
  },
  "Rechnung": {
    "count": 7,
    "lastSeenTime": 1586879558728,
    "difficulty": 0.975609756097561,
    "difficultyWeight": 41
  },
  "Panzerfaust": {
    "count": 7,
    "lastSeenTime": 1586955125492
  },
  "Wahl": {
    "count": 4,
    "lastSeenTime": 1586873925511
  },
  "Farbe": {
    "count": 6,
    "lastSeenTime": 1586916266718
  },
  "Hafenbecken": {
    "count": 7,
    "lastSeenTime": 1586958264090
  },
  "Mario": {
    "count": 6,
    "lastSeenTime": 1586956599160
  },
  "Zahnstocher": {
    "count": 4,
    "lastSeenTime": 1586900670423
  },
  "Mumie": {
    "count": 4,
    "lastSeenTime": 1586920005337,
    "difficulty": 1,
    "difficultyWeight": 26
  },
  "Stativ": {
    "count": 4,
    "lastSeenTime": 1586900162376
  },
  "Fernglas": {
    "count": 10,
    "lastSeenTime": 1586921022722
  },
  "parallel": {
    "count": 6,
    "lastSeenTime": 1586916565022,
    "difficulty": 1,
    "difficultyWeight": 3
  },
  "Rüstung": {
    "count": 7,
    "lastSeenTime": 1586827071400
  },
  "Gang": {
    "count": 4,
    "lastSeenTime": 1586813707177
  },
  "Insel": {
    "count": 8,
    "lastSeenTime": 1586959020405,
    "difficulty": 0.631578947368421,
    "difficultyWeight": 19
  },
  "Versicherungskaufmann": {
    "count": 8,
    "lastSeenTime": 1586955088291
  },
  "Euro": {
    "count": 7,
    "lastSeenTime": 1586959850831
  },
  "Charlie Chaplin": {
    "count": 12,
    "lastSeenTime": 1586958907689,
    "publicGameCount": 1
  },
  "Barbecue": {
    "count": 8,
    "lastSeenTime": 1586914554202,
    "difficulty": 0.8,
    "difficultyWeight": 20
  },
  "Haus": {
    "count": 4,
    "lastSeenTime": 1586785382597,
    "difficulty": 0.5,
    "difficultyWeight": 10
  },
  "Nachtclub": {
    "count": 6,
    "lastSeenTime": 1586881045754
  },
  "Toaster": {
    "count": 9,
    "lastSeenTime": 1586896242674
  },
  "Nacht": {
    "count": 10,
    "lastSeenTime": 1586958769937
  },
  "Elmo": {
    "count": 8,
    "lastSeenTime": 1586955620355
  },
  "König der Löwen": {
    "count": 6,
    "lastSeenTime": 1586895858023,
    "difficulty": 1,
    "difficultyWeight": 5
  },
  "Mona Lisa": {
    "count": 5,
    "lastSeenTime": 1586920751201
  },
  "Muskete": {
    "count": 5,
    "lastSeenTime": 1586956033779,
    "difficulty": 0,
    "difficultyWeight": 2
  },
  "Handgelenk": {
    "count": 3,
    "lastSeenTime": 1586957888650,
    "publicGameCount": 1,
    "difficulty": 1,
    "difficultyWeight": 13
  },
  "Lasagne": {
    "count": 9,
    "lastSeenTime": 1586908494479
  },
  "KFC": {
    "count": 7,
    "lastSeenTime": 1586909250467
  },
  "Überlebender": {
    "count": 10,
    "lastSeenTime": 1586868863888
  },
  "Bandnudel": {
    "count": 5,
    "lastSeenTime": 1586879785134
  },
  "Lemur": {
    "count": 10,
    "lastSeenTime": 1586956453935
  },
  "Walross": {
    "count": 7,
    "lastSeenTime": 1586956197682,
    "difficulty": 0.7142857142857143,
    "difficultyWeight": 7
  },
  "Baseballschläger": {
    "count": 10,
    "lastSeenTime": 1586891023780
  },
  "Oktopus": {
    "count": 12,
    "lastSeenTime": 1586957837428,
    "difficulty": 1,
    "difficultyWeight": 6
  },
  "Rennen": {
    "count": 6,
    "lastSeenTime": 1586915767461,
    "difficulty": 0.5,
    "difficultyWeight": 10
  },
  "Mikrofon": {
    "count": 8,
    "lastSeenTime": 1586895872235,
    "publicGameCount": 1,
    "difficulty": 0.6,
    "difficultyWeight": 10
  },
  "Bogen": {
    "count": 6,
    "lastSeenTime": 1586806308278
  },
  "Superman": {
    "count": 9,
    "lastSeenTime": 1586920336794,
    "publicGameCount": 1,
    "difficulty": 0.5555555555555556,
    "difficultyWeight": 9
  },
  "Dudelsack": {
    "count": 7,
    "lastSeenTime": 1586914181975
  },
  "Lippen": {
    "count": 12,
    "lastSeenTime": 1586903726114,
    "difficulty": 1,
    "difficultyWeight": 9
  },
  "Fotograf": {
    "count": 14,
    "lastSeenTime": 1586958561895,
    "difficulty": 1,
    "difficultyWeight": 18
  },
  "Popcorn": {
    "count": 8,
    "lastSeenTime": 1586955809521,
    "publicGameCount": 1,
    "difficulty": 0.3333333333333333,
    "difficultyWeight": 6
  },
  "Haarschuppen": {
    "count": 5,
    "lastSeenTime": 1586896700238
  },
  "vollständig": {
    "count": 8,
    "lastSeenTime": 1586958547642
  },
  "Batman": {
    "count": 5,
    "lastSeenTime": 1586914243836
  },
  "Vulkanologe": {
    "count": 7,
    "lastSeenTime": 1586956287941
  },
  "Geschirrschrank": {
    "count": 3,
    "lastSeenTime": 1586782917680
  },
  "Taschentuch": {
    "count": 4,
    "lastSeenTime": 1586879263241
  },
  "Lamm": {
    "count": 8,
    "lastSeenTime": 1586955300241,
    "difficulty": 1,
    "difficultyWeight": 1
  },
  "Megafon": {
    "count": 7,
    "lastSeenTime": 1586920651904
  },
  "Lockenwickler": {
    "count": 4,
    "lastSeenTime": 1586784602447
  },
  "Bienenstich": {
    "count": 10,
    "lastSeenTime": 1586916368116,
    "publicGameCount": 1,
    "difficulty": 0.4117647058823529,
    "difficultyWeight": 17
  },
  "Ohrring": {
    "count": 3,
    "lastSeenTime": 1586784746074
  },
  "wütend": {
    "count": 8,
    "lastSeenTime": 1586920029620
  },
  "Hölle": {
    "count": 4,
    "lastSeenTime": 1586901682745,
    "difficulty": 1,
    "difficultyWeight": 11
  },
  "Plätzchen": {
    "count": 4,
    "lastSeenTime": 1586827227421
  },
  "Traktor": {
    "count": 11,
    "lastSeenTime": 1586908787449
  },
  "Skorpion": {
    "count": 7,
    "lastSeenTime": 1586959633128
  },
  "Fell": {
    "count": 6,
    "lastSeenTime": 1586955517664
  },
  "Lady Gaga": {
    "count": 5,
    "lastSeenTime": 1586901511433
  },
  "übergewichtig": {
    "count": 6,
    "lastSeenTime": 1586959117095
  },
  "Himmelbett": {
    "count": 9,
    "lastSeenTime": 1586914903455
  },
  "Boot": {
    "count": 7,
    "lastSeenTime": 1586958683104,
    "difficulty": 0.5625,
    "difficultyWeight": 16,
    "publicGameCount": 1
  },
  "Chinesische Mauer": {
    "count": 4,
    "lastSeenTime": 1586890378764
  },
  "Link": {
    "count": 5,
    "lastSeenTime": 1586907863533
  },
  "Cartoon": {
    "count": 5,
    "lastSeenTime": 1586907707534,
    "publicGameCount": 1,
    "difficulty": 1,
    "difficultyWeight": 10
  },
  "Seegurke": {
    "count": 7,
    "lastSeenTime": 1586897025198,
    "difficulty": 1,
    "difficultyWeight": 7
  },
  "Zucker": {
    "count": 5,
    "lastSeenTime": 1586907778315
  },
  "Bambus": {
    "count": 6,
    "lastSeenTime": 1586954617410
  },
  "Marathon": {
    "count": 7,
    "lastSeenTime": 1586955392320
  },
  "Donner": {
    "count": 5,
    "lastSeenTime": 1586832635442
  },
  "Schottland": {
    "count": 5,
    "lastSeenTime": 1586957863554
  },
  "beißen": {
    "count": 6,
    "lastSeenTime": 1586889525947
  },
  "Konzert": {
    "count": 7,
    "lastSeenTime": 1586914403534
  },
  "Vogel": {
    "count": 2,
    "lastSeenTime": 1586896941227,
    "difficulty": 0.86,
    "difficultyWeight": 50
  },
  "Maniküre": {
    "count": 9,
    "lastSeenTime": 1586957960523
  },
  "Personenschützer": {
    "count": 5,
    "lastSeenTime": 1586920094776
  },
  "Nagel": {
    "count": 12,
    "lastSeenTime": 1586955456078
  },
  "Flammkuchen": {
    "count": 5,
    "lastSeenTime": 1586896242674
  },
  "Kopftuch": {
    "count": 5,
    "lastSeenTime": 1586915949972,
    "difficulty": 0.8571428571428571,
    "difficultyWeight": 14
  },
  "Faustkampf": {
    "count": 4,
    "lastSeenTime": 1586816692499
  },
  "Bikini": {
    "count": 7,
    "lastSeenTime": 1586900635930
  },
  "schütteln": {
    "count": 6,
    "lastSeenTime": 1586955334605
  },
  "Klebestift": {
    "count": 6,
    "lastSeenTime": 1586957950175,
    "difficulty": 1,
    "difficultyWeight": 19
  },
  "Florist": {
    "count": 6,
    "lastSeenTime": 1586959522880
  },
  "Ananas": {
    "count": 6,
    "lastSeenTime": 1586826228383
  },
  "Oper": {
    "count": 6,
    "lastSeenTime": 1586916574944,
    "difficulty": 0.5714285714285714,
    "difficultyWeight": 7
  },
  "Klippe": {
    "count": 5,
    "lastSeenTime": 1586900422423
  },
  "Riese": {
    "count": 5,
    "lastSeenTime": 1586858562011,
    "difficulty": 0.5,
    "difficultyWeight": 8
  },
  "Krankenhaus": {
    "count": 7,
    "lastSeenTime": 1586908697152
  },
  "müde": {
    "count": 2,
    "lastSeenTime": 1586958171361
  },
  "prokrastinieren": {
    "count": 3,
    "lastSeenTime": 1586806544125
  },
  "Wrestler": {
    "count": 6,
    "lastSeenTime": 1586889892808
  },
  "Eiskaffee": {
    "count": 7,
    "lastSeenTime": 1586914270411
  },
  "Auspuff": {
    "count": 7,
    "lastSeenTime": 1586954925078
  },
  "erinnern": {
    "count": 3,
    "lastSeenTime": 1586909119149
  },
  "Walnuss": {
    "count": 4,
    "lastSeenTime": 1586903726114,
    "difficulty": 0.875,
    "difficultyWeight": 16
  },
  "Pech": {
    "count": 8,
    "lastSeenTime": 1586899941163,
    "difficulty": 1,
    "difficultyWeight": 2
  },
  "Garage": {
    "count": 8,
    "lastSeenTime": 1586915200929
  },
  "Lego": {
    "count": 6,
    "lastSeenTime": 1586838924837,
    "difficulty": 0.6,
    "difficultyWeight": 5
  },
  "Ballon": {
    "count": 6,
    "lastSeenTime": 1586908541127,
    "difficulty": 0.72,
    "difficultyWeight": 50
  },
  "Barkeeper": {
    "count": 11,
    "lastSeenTime": 1586958809891
  },
  "vergrößern": {
    "count": 5,
    "lastSeenTime": 1586913973573
  },
  "insolvent": {
    "count": 9,
    "lastSeenTime": 1586915806469
  },
  "Demonstration": {
    "count": 4,
    "lastSeenTime": 1586859897712,
    "publicGameCount": 1,
    "difficulty": 0.375,
    "difficultyWeight": 8
  },
  "denken": {
    "count": 4,
    "lastSeenTime": 1586838914729
  },
  "Goldtopf": {
    "count": 7,
    "lastSeenTime": 1586920275708
  },
  "Wanderstock": {
    "count": 9,
    "lastSeenTime": 1586915417224,
    "difficulty": 1,
    "difficultyWeight": 20
  },
  "Papagei": {
    "count": 8,
    "lastSeenTime": 1586916141879
  },
  "Fledermaus": {
    "count": 6,
    "lastSeenTime": 1586890056241,
    "difficulty": 1,
    "difficultyWeight": 6
  },
  "Sudoku": {
    "count": 1,
    "lastSeenTime": 1586552881359,
    "difficulty": 1,
    "difficultyWeight": 1
  },
  "Cola": {
    "count": 11,
    "lastSeenTime": 1586958983867,
    "difficulty": 1,
    "difficultyWeight": 8
  },
  "Kakao": {
    "count": 5,
    "lastSeenTime": 1586904343119,
    "publicGameCount": 1
  },
  "untergewichtig": {
    "count": 3,
    "lastSeenTime": 1586914491372,
    "difficulty": 0.8421052631578947,
    "difficultyWeight": 19
  },
  "Zeitlupe": {
    "count": 9,
    "lastSeenTime": 1586920053882,
    "difficulty": 0.8888888888888888,
    "difficultyWeight": 36
  },
  "Österreich": {
    "count": 9,
    "lastSeenTime": 1586915919352
  },
  "Nintendo": {
    "count": 4,
    "lastSeenTime": 1586913633876,
    "difficulty": 1,
    "difficultyWeight": 3
  },
  "Eltern": {
    "count": 5,
    "lastSeenTime": 1586955429782,
    "difficulty": 0.4666666666666667,
    "difficultyWeight": 15
  },
  "Tannenbaum": {
    "count": 9,
    "lastSeenTime": 1586907472056,
    "difficulty": 1,
    "difficultyWeight": 5
  },
  "Hyäne": {
    "count": 7,
    "lastSeenTime": 1586806967390,
    "difficulty": 1,
    "difficultyWeight": 4
  },
  "Komödie": {
    "count": 4,
    "lastSeenTime": 1586589955120
  },
  "Daune": {
    "count": 8,
    "lastSeenTime": 1586914806527
  },
  "Axt": {
    "count": 6,
    "lastSeenTime": 1586901361357
  },
  "Heuschrecke": {
    "count": 5,
    "lastSeenTime": 1586908118744,
    "difficulty": 1,
    "difficultyWeight": 25
  },
  "Baumkrone": {
    "count": 8,
    "lastSeenTime": 1586881235651
  },
  "Bier": {
    "count": 7,
    "lastSeenTime": 1586921059401,
    "publicGameCount": 1,
    "difficulty": 0.967741935483871,
    "difficultyWeight": 93
  },
  "Maibaum": {
    "count": 8,
    "lastSeenTime": 1586915299932
  },
  "Haarfarbe": {
    "count": 7,
    "lastSeenTime": 1586959814583,
    "difficulty": 1,
    "difficultyWeight": 43
  },
  "Webseite": {
    "count": 5,
    "lastSeenTime": 1586815794279,
    "difficulty": 0.6923076923076923,
    "difficultyWeight": 26
  },
  "Gift": {
    "count": 5,
    "lastSeenTime": 1586959020405
  },
  "Buch": {
    "count": 5,
    "lastSeenTime": 1586956009184
  },
  "Pflug": {
    "count": 4,
    "lastSeenTime": 1586901375521,
    "difficulty": 1,
    "difficultyWeight": 4
  },
  "Blut": {
    "count": 7,
    "lastSeenTime": 1586958983867
  },
  "Igel": {
    "count": 9,
    "lastSeenTime": 1586859619922
  },
  "machomäßig": {
    "count": 3,
    "lastSeenTime": 1586805850042,
    "difficulty": 0.9,
    "difficultyWeight": 10
  },
  "Schnürsenkel": {
    "count": 9,
    "lastSeenTime": 1586920604951,
    "difficulty": 0.42857142857142855,
    "difficultyWeight": 7
  },
  "Felsen": {
    "count": 12,
    "lastSeenTime": 1586954834961,
    "difficulty": 0.5882352941176471,
    "difficultyWeight": 17,
    "publicGameCount": 1
  },
  "Mähdrescher": {
    "count": 5,
    "lastSeenTime": 1586901067015
  },
  "Anglerfisch": {
    "count": 5,
    "lastSeenTime": 1586785741382
  },
  "Kanarienvogel": {
    "count": 3,
    "lastSeenTime": 1586873649256,
    "difficulty": 1,
    "difficultyWeight": 10
  },
  "Lautstärke": {
    "count": 8,
    "lastSeenTime": 1586890045700
  },
  "Dynamo": {
    "count": 5,
    "lastSeenTime": 1586869422590
  },
  "Brust": {
    "count": 5,
    "lastSeenTime": 1586897420912,
    "difficulty": 1,
    "difficultyWeight": 1
  },
  "Fahrer": {
    "count": 5,
    "lastSeenTime": 1586958561895
  },
  "Schaukelpferd": {
    "count": 5,
    "lastSeenTime": 1586958697329,
    "difficulty": 1,
    "difficultyWeight": 5
  },
  "Schatten": {
    "count": 5,
    "lastSeenTime": 1586840034942,
    "difficulty": 1,
    "difficultyWeight": 5
  },
  "Entensuppe": {
    "count": 6,
    "lastSeenTime": 1586897164255,
    "difficulty": 1,
    "difficultyWeight": 6
  },
  "Abraham Lincoln": {
    "count": 6,
    "lastSeenTime": 1586956599160
  },
  "Papageientaucher": {
    "count": 11,
    "lastSeenTime": 1586956363582
  },
  "Milch": {
    "count": 4,
    "lastSeenTime": 1586959193290,
    "difficulty": 0.9166666666666666,
    "difficultyWeight": 12
  },
  "Yeti": {
    "count": 5,
    "lastSeenTime": 1586955152448
  },
  "dreckig": {
    "count": 4,
    "lastSeenTime": 1586920552243
  },
  "Steigung": {
    "count": 4,
    "lastSeenTime": 1586908838335
  },
  "Schaumschläger": {
    "count": 10,
    "lastSeenTime": 1586955681359,
    "difficulty": 1,
    "difficultyWeight": 9
  },
  "Möbel": {
    "count": 5,
    "lastSeenTime": 1586909133282,
    "difficulty": 1,
    "difficultyWeight": 27
  },
  "Feueralarm": {
    "count": 4,
    "lastSeenTime": 1586914428482
  },
  "Fahrkarte": {
    "count": 9,
    "lastSeenTime": 1586904061376
  },
  "Kredit": {
    "count": 7,
    "lastSeenTime": 1586920972036
  },
  "Alphorn": {
    "count": 10,
    "lastSeenTime": 1586958278719
  },
  "Partnerlook": {
    "count": 8,
    "lastSeenTime": 1586916352907,
    "difficulty": 1,
    "difficultyWeight": 9
  },
  "Panflöte": {
    "count": 5,
    "lastSeenTime": 1586827260902
  },
  "Blumenkohl": {
    "count": 5,
    "lastSeenTime": 1586958809891
  },
  "Geist": {
    "count": 3,
    "lastSeenTime": 1586868785173,
    "difficulty": 0.75,
    "difficultyWeight": 32
  },
  "Nordpol": {
    "count": 5,
    "lastSeenTime": 1586958360375,
    "publicGameCount": 1,
    "difficulty": 0.8,
    "difficultyWeight": 10
  },
  "Fanta": {
    "count": 8,
    "lastSeenTime": 1586919783421,
    "difficulty": 0.6666666666666666,
    "difficultyWeight": 18
  },
  "Periskop": {
    "count": 6,
    "lastSeenTime": 1586915731157
  },
  "erbrechen": {
    "count": 6,
    "lastSeenTime": 1586954798444
  },
  "Bonbon": {
    "count": 9,
    "lastSeenTime": 1586915703672,
    "difficulty": 0.625,
    "difficultyWeight": 8,
    "publicGameCount": 1
  },
  "Oboe": {
    "count": 9,
    "lastSeenTime": 1586916574944
  },
  "Schnurrbart": {
    "count": 8,
    "lastSeenTime": 1586914157190,
    "publicGameCount": 1,
    "difficulty": 0.7142857142857143,
    "difficultyWeight": 7
  },
  "Höhlenmensch": {
    "count": 7,
    "lastSeenTime": 1586954975722
  },
  "Büroklammer": {
    "count": 13,
    "lastSeenTime": 1586959570753,
    "difficulty": 0.7608695652173914,
    "difficultyWeight": 46
  },
  "Tunnelblick": {
    "count": 6,
    "lastSeenTime": 1586958077398
  },
  "Besenstiel": {
    "count": 3,
    "lastSeenTime": 1586903648591,
    "difficulty": 1,
    "difficultyWeight": 27
  },
  "Kreuzung": {
    "count": 8,
    "lastSeenTime": 1586913670980,
    "difficulty": 0.9423076923076923,
    "difficultyWeight": 52
  },
  "Zentaur": {
    "count": 4,
    "lastSeenTime": 1586908583500,
    "difficulty": 0.8,
    "difficultyWeight": 5
  },
  "Papaya": {
    "count": 6,
    "lastSeenTime": 1586955263523
  },
  "Notch": {
    "count": 4,
    "lastSeenTime": 1586826344267
  },
  "Arm": {
    "count": 9,
    "lastSeenTime": 1586828011219,
    "difficulty": 1,
    "difficultyWeight": 6
  },
  "Dolch": {
    "count": 8,
    "lastSeenTime": 1586914565215
  },
  "DNS": {
    "count": 4,
    "lastSeenTime": 1586873397605,
    "difficulty": 1,
    "difficultyWeight": 2
  },
  "Seelöwe": {
    "count": 5,
    "lastSeenTime": 1586858858134,
    "difficulty": 1,
    "difficultyWeight": 6
  },
  "Blutspende": {
    "count": 8,
    "lastSeenTime": 1586955113604,
    "publicGameCount": 1
  },
  "Astronaut": {
    "count": 6,
    "lastSeenTime": 1586959279701
  },
  "Frankreich": {
    "count": 3,
    "lastSeenTime": 1586596653489
  },
  "Beine": {
    "count": 8,
    "lastSeenTime": 1586913768003
  },
  "Logo": {
    "count": 7,
    "lastSeenTime": 1586920015464
  },
  "Freitag": {
    "count": 4,
    "lastSeenTime": 1586812381709,
    "publicGameCount": 1,
    "difficulty": 0.6875,
    "difficultyWeight": 16
  },
  "Süßholz raspeln": {
    "count": 4,
    "lastSeenTime": 1586812346307
  },
  "Schock": {
    "count": 4,
    "lastSeenTime": 1586897287249
  },
  "Alufolie": {
    "count": 8,
    "lastSeenTime": 1586900724985,
    "difficulty": 1,
    "difficultyWeight": 6
  },
  "Geysir": {
    "count": 3,
    "lastSeenTime": 1586727718047
  },
  "Kirche": {
    "count": 4,
    "lastSeenTime": 1586651112698,
    "difficulty": 0.5454545454545454,
    "difficultyWeight": 11
  },
  "Kinderwagen": {
    "count": 4,
    "lastSeenTime": 1586816753545
  },
  "Kuckucksuhr": {
    "count": 4,
    "lastSeenTime": 1586815755534
  },
  "Grinch": {
    "count": 10,
    "lastSeenTime": 1586958239433,
    "difficulty": 1,
    "difficultyWeight": 10
  },
  "Klempner": {
    "count": 6,
    "lastSeenTime": 1586903531503,
    "publicGameCount": 1,
    "difficulty": 0.75,
    "difficultyWeight": 8
  },
  "Geröll": {
    "count": 7,
    "lastSeenTime": 1586903634298
  },
  "Hologramm": {
    "count": 4,
    "lastSeenTime": 1586726805561
  },
  "Rollladen": {
    "count": 4,
    "lastSeenTime": 1586839583890
  },
  "Dikdik": {
    "count": 4,
    "lastSeenTime": 1586890005207
  },
  "Gladiator": {
    "count": 5,
    "lastSeenTime": 1586921032863,
    "difficulty": 0.8421052631578947,
    "difficultyWeight": 19,
    "publicGameCount": 1
  },
  "Brasilien": {
    "count": 6,
    "lastSeenTime": 1586827712672
  },
  "Glanz": {
    "count": 4,
    "lastSeenTime": 1586913853735
  },
  "Karte": {
    "count": 6,
    "lastSeenTime": 1586919287618,
    "difficulty": 1,
    "difficultyWeight": 15
  },
  "Silber": {
    "count": 6,
    "lastSeenTime": 1586896166029,
    "publicGameCount": 1
  },
  "Sicherheitsdienst": {
    "count": 5,
    "lastSeenTime": 1586915949972
  },
  "Uranus": {
    "count": 5,
    "lastSeenTime": 1586874623611,
    "difficulty": 1,
    "difficultyWeight": 11
  },
  "Knöchel": {
    "count": 9,
    "lastSeenTime": 1586958944806
  },
  "Eiche": {
    "count": 4,
    "lastSeenTime": 1586920566543,
    "difficulty": 1,
    "difficultyWeight": 79
  },
  "Schnecke": {
    "count": 9,
    "lastSeenTime": 1586919773316
  },
  "Bleistift": {
    "count": 3,
    "lastSeenTime": 1586717347143
  },
  "Reh": {
    "count": 5,
    "lastSeenTime": 1586914831496,
    "publicGameCount": 1,
    "difficulty": 0.6,
    "difficultyWeight": 5
  },
  "Anzug": {
    "count": 3,
    "lastSeenTime": 1586806231030,
    "difficulty": 1,
    "difficultyWeight": 5
  },
  "Süden": {
    "count": 5,
    "lastSeenTime": 1586782502910,
    "difficulty": 0.2222222222222222,
    "difficultyWeight": 9
  },
  "Spitzer": {
    "count": 5,
    "lastSeenTime": 1586958375030
  },
  "lesen": {
    "count": 5,
    "lastSeenTime": 1586958973713,
    "difficulty": 0.8387096774193549,
    "difficultyWeight": 31
  },
  "Köder": {
    "count": 3,
    "lastSeenTime": 1586832725500,
    "difficulty": 1,
    "difficultyWeight": 1
  },
  "Nike": {
    "count": 9,
    "lastSeenTime": 1586959492934,
    "difficulty": 0.7857142857142857,
    "difficultyWeight": 14
  },
  "Vault boy": {
    "count": 3,
    "lastSeenTime": 1586907877908
  },
  "Handy": {
    "count": 6,
    "lastSeenTime": 1586890188914,
    "difficulty": 1,
    "difficultyWeight": 1
  },
  "Garten": {
    "count": 6,
    "lastSeenTime": 1586901605672
  },
  "Fuchs": {
    "count": 3,
    "lastSeenTime": 1586920566543,
    "difficulty": 0.8666666666666667,
    "difficultyWeight": 30
  },
  "Musik": {
    "count": 3,
    "lastSeenTime": 1586879877998,
    "difficulty": 1,
    "difficultyWeight": 6
  },
  "niedrig": {
    "count": 8,
    "lastSeenTime": 1586880759383
  },
  "Scheibenwischer": {
    "count": 9,
    "lastSeenTime": 1586959304612
  },
  "Roboter": {
    "count": 6,
    "lastSeenTime": 1586909133282
  },
  "Säge": {
    "count": 2,
    "lastSeenTime": 1586633399505,
    "difficulty": 0.4,
    "difficultyWeight": 10
  },
  "blau": {
    "count": 9,
    "lastSeenTime": 1586914403534,
    "publicGameCount": 1,
    "difficulty": 0.96875,
    "difficultyWeight": 32
  },
  "Treibsand": {
    "count": 7,
    "lastSeenTime": 1586959695898
  },
  "Segelflieger": {
    "count": 5,
    "lastSeenTime": 1586954774095
  },
  "Zahnseide": {
    "count": 4,
    "lastSeenTime": 1586903333845
  },
  "glühen": {
    "count": 5,
    "lastSeenTime": 1586818299882
  },
  "Inuit": {
    "count": 10,
    "lastSeenTime": 1586914258463
  },
  "Flaschenpost": {
    "count": 6,
    "lastSeenTime": 1586907451820
  },
  "Christbaumkugel": {
    "count": 8,
    "lastSeenTime": 1586957903080
  },
  "Hummer": {
    "count": 12,
    "lastSeenTime": 1586958052113
  },
  "Vin Diesel": {
    "count": 10,
    "lastSeenTime": 1586915186679
  },
  "offen": {
    "count": 6,
    "lastSeenTime": 1586889697457
  },
  "Windmühle": {
    "count": 3,
    "lastSeenTime": 1586833318970,
    "difficulty": 0.9166666666666666,
    "difficultyWeight": 24
  },
  "fernsehen": {
    "count": 5,
    "lastSeenTime": 1586915846579
  },
  "Zeitmaschine": {
    "count": 7,
    "lastSeenTime": 1586897211835
  },
  "Tasche": {
    "count": 5,
    "lastSeenTime": 1586858696188,
    "publicGameCount": 1,
    "difficulty": 1,
    "difficultyWeight": 10
  },
  "Abendessen": {
    "count": 8,
    "lastSeenTime": 1586903965187,
    "difficulty": 0.75,
    "difficultyWeight": 8
  },
  "Moskito": {
    "count": 11,
    "lastSeenTime": 1586959714261,
    "difficulty": 0.5714285714285714,
    "difficultyWeight": 7
  },
  "Lava": {
    "count": 8,
    "lastSeenTime": 1586902766371,
    "difficulty": 1,
    "difficultyWeight": 6
  },
  "Eingabestift": {
    "count": 5,
    "lastSeenTime": 1586915021163
  },
  "Tropfen": {
    "count": 4,
    "lastSeenTime": 1586958748793,
    "difficulty": 0.875,
    "difficultyWeight": 40
  },
  "Zoowärter": {
    "count": 8,
    "lastSeenTime": 1586959290196,
    "difficulty": 1,
    "difficultyWeight": 3
  },
  "Flugzeug": {
    "count": 4,
    "lastSeenTime": 1586915149693
  },
  "doppelt": {
    "count": 4,
    "lastSeenTime": 1586806127412
  },
  "Senf": {
    "count": 5,
    "lastSeenTime": 1586900422423
  },
  "Klassenzimmer": {
    "count": 5,
    "lastSeenTime": 1586908801649,
    "difficulty": 0.75,
    "difficultyWeight": 4
  },
  "schminken": {
    "count": 7,
    "lastSeenTime": 1586955209647
  },
  "Knoten": {
    "count": 6,
    "lastSeenTime": 1586909005772
  },
  "kahl": {
    "count": 7,
    "lastSeenTime": 1586957903080
  },
  "Ketchup": {
    "count": 6,
    "lastSeenTime": 1586875034252,
    "publicGameCount": 1,
    "difficulty": 0.5555555555555556,
    "difficultyWeight": 9
  },
  "Armbrust": {
    "count": 5,
    "lastSeenTime": 1586897226069
  },
  "Kehle": {
    "count": 5,
    "lastSeenTime": 1586900524596
  },
  "Deckenventilator": {
    "count": 7,
    "lastSeenTime": 1586815355269
  },
  "Hobbit": {
    "count": 2,
    "lastSeenTime": 1586784590637,
    "difficulty": 1,
    "difficultyWeight": 18
  },
  "Boxer": {
    "count": 5,
    "lastSeenTime": 1586814868590,
    "difficulty": 1,
    "difficultyWeight": 1
  },
  "Lupe": {
    "count": 9,
    "lastSeenTime": 1586897078678,
    "difficulty": 1,
    "difficultyWeight": 24
  },
  "Steinbock": {
    "count": 8,
    "lastSeenTime": 1586914631184,
    "difficulty": 0.9696969696969697,
    "difficultyWeight": 33
  },
  "Tor": {
    "count": 7,
    "lastSeenTime": 1586958636395,
    "difficulty": 1,
    "difficultyWeight": 10
  },
  "Kneipe": {
    "count": 10,
    "lastSeenTime": 1586890339661,
    "difficulty": 0.8,
    "difficultyWeight": 15
  },
  "Zement": {
    "count": 7,
    "lastSeenTime": 1586895733096
  },
  "Träne": {
    "count": 8,
    "lastSeenTime": 1586873585720,
    "difficulty": 0.7222222222222222,
    "difficultyWeight": 18
  },
  "Veranda": {
    "count": 7,
    "lastSeenTime": 1586954727368
  },
  "Oberarm": {
    "count": 3,
    "lastSeenTime": 1586806833684
  },
  "Marmelade": {
    "count": 7,
    "lastSeenTime": 1586903836685,
    "difficulty": 1,
    "difficultyWeight": 40
  },
  "Banjo": {
    "count": 8,
    "lastSeenTime": 1586958200965
  },
  "Herz": {
    "count": 7,
    "lastSeenTime": 1586879684720,
    "difficulty": 0.8695652173913043,
    "difficultyWeight": 23,
    "publicGameCount": 1
  },
  "Vision": {
    "count": 10,
    "lastSeenTime": 1586907654631
  },
  "Katze": {
    "count": 7,
    "lastSeenTime": 1586874013016
  },
  "Knopf": {
    "count": 2,
    "lastSeenTime": 1586710517071,
    "difficulty": 0.4166666666666667,
    "difficultyWeight": 12
  },
  "biegen": {
    "count": 7,
    "lastSeenTime": 1586916455715
  },
  "Prüfung": {
    "count": 9,
    "lastSeenTime": 1586955948209
  },
  "Psychologe": {
    "count": 6,
    "lastSeenTime": 1586959585641
  },
  "zusammenzucken": {
    "count": 10,
    "lastSeenTime": 1586915354278
  },
  "Baumhaus": {
    "count": 5,
    "lastSeenTime": 1586909260571,
    "difficulty": 1,
    "difficultyWeight": 5
  },
  "Alpaka": {
    "count": 12,
    "lastSeenTime": 1586900125985
  },
  "Hieroglyphen": {
    "count": 12,
    "lastSeenTime": 1586958596468
  },
  "Saugglocke": {
    "count": 7,
    "lastSeenTime": 1586832373335
  },
  "Kasino": {
    "count": 8,
    "lastSeenTime": 1586915974465
  },
  "Funke": {
    "count": 7,
    "lastSeenTime": 1586955385136,
    "difficulty": 1,
    "difficultyWeight": 10
  },
  "Venus": {
    "count": 10,
    "lastSeenTime": 1586959076238,
    "difficulty": 0.7,
    "difficultyWeight": 20
  },
  "Feuersalamander": {
    "count": 1,
    "lastSeenTime": 1586561707947
  },
  "Lebensmittel": {
    "count": 3,
    "lastSeenTime": 1586726805561,
    "difficulty": 1,
    "difficultyWeight": 3
  },
  "Säule": {
    "count": 4,
    "lastSeenTime": 1586895477389,
    "difficulty": 0.6666666666666666,
    "difficultyWeight": 6
  },
  "dürr": {
    "count": 3,
    "lastSeenTime": 1586955913834
  },
  "Audi": {
    "count": 6,
    "lastSeenTime": 1586880818922,
    "difficulty": 1,
    "difficultyWeight": 6
  },
  "langsam": {
    "count": 3,
    "lastSeenTime": 1586709182856
  },
  "Lutscher": {
    "count": 7,
    "lastSeenTime": 1586869733694
  },
  "haarig": {
    "count": 5,
    "lastSeenTime": 1586889811471,
    "publicGameCount": 1,
    "difficulty": 0.6666666666666666,
    "difficultyWeight": 15
  },
  "Monobraue": {
    "count": 9,
    "lastSeenTime": 1586958697329
  },
  "bekifft": {
    "count": 10,
    "lastSeenTime": 1586915716336
  },
  "Elfenbein": {
    "count": 11,
    "lastSeenTime": 1586959090713
  },
  "Medusa": {
    "count": 4,
    "lastSeenTime": 1586832990793
  },
  "Schlumpf": {
    "count": 6,
    "lastSeenTime": 1586916164227,
    "publicGameCount": 1
  },
  "Muschel": {
    "count": 16,
    "lastSeenTime": 1586959232575,
    "publicGameCount": 1,
    "difficulty": 0.3333333333333333,
    "difficultyWeight": 6
  },
  "Kartoffel": {
    "count": 4,
    "lastSeenTime": 1586901839909
  },
  "Maler": {
    "count": 7,
    "lastSeenTime": 1586908838335
  },
  "stoßen": {
    "count": 5,
    "lastSeenTime": 1586901166652
  },
  "Nemo": {
    "count": 4,
    "lastSeenTime": 1586954910317,
    "publicGameCount": 1,
    "difficulty": 1,
    "difficultyWeight": 15
  },
  "Schimpanse": {
    "count": 6,
    "lastSeenTime": 1586879248947,
    "difficulty": 0.7777777777777778,
    "difficultyWeight": 9
  },
  "schlafen": {
    "count": 7,
    "lastSeenTime": 1586956048018,
    "difficulty": 0.4,
    "difficultyWeight": 5
  },
  "Einzelhandel": {
    "count": 5,
    "lastSeenTime": 1586879799373,
    "publicGameCount": 1
  },
  "Frühling": {
    "count": 8,
    "lastSeenTime": 1586957837428,
    "difficulty": 0.3333333333333333,
    "difficultyWeight": 6
  },
  "Corn Flakes": {
    "count": 5,
    "lastSeenTime": 1586919687264,
    "difficulty": 0.875,
    "difficultyWeight": 8
  },
  "Tarzan": {
    "count": 7,
    "lastSeenTime": 1586959455308
  },
  "Baumstumpf": {
    "count": 7,
    "lastSeenTime": 1586902713347
  },
  "schneiden": {
    "count": 7,
    "lastSeenTime": 1586919348003
  },
  "Kopfhörer": {
    "count": 7,
    "lastSeenTime": 1586955609964,
    "difficulty": 1,
    "difficultyWeight": 6
  },
  "Suppenkelle": {
    "count": 3,
    "lastSeenTime": 1586959482662
  },
  "Paralleluniversum": {
    "count": 5,
    "lastSeenTime": 1586826400303
  },
  "Friedhof": {
    "count": 8,
    "lastSeenTime": 1586954627819,
    "difficulty": 1,
    "difficultyWeight": 6
  },
  "Auster": {
    "count": 5,
    "lastSeenTime": 1586958100659
  },
  "Flammenwerfer": {
    "count": 6,
    "lastSeenTime": 1586956599160
  },
  "Bürgersteig": {
    "count": 4,
    "lastSeenTime": 1586913482555
  },
  "Luxemburg": {
    "count": 5,
    "lastSeenTime": 1586890217261
  },
  "Hotel": {
    "count": 6,
    "lastSeenTime": 1586816892701
  },
  "Spirale": {
    "count": 3,
    "lastSeenTime": 1586706667073
  },
  "schauen": {
    "count": 7,
    "lastSeenTime": 1586879469524
  },
  "Lavalampe": {
    "count": 3,
    "lastSeenTime": 1586807857048,
    "difficulty": 1,
    "difficultyWeight": 3
  },
  "Nussschale": {
    "count": 6,
    "lastSeenTime": 1586895833706
  },
  "Kette": {
    "count": 6,
    "lastSeenTime": 1586956048018
  },
  "Trinkgeld": {
    "count": 4,
    "lastSeenTime": 1586868835247
  },
  "Leuchtturm": {
    "count": 6,
    "lastSeenTime": 1586896740857
  },
  "Fingernagel": {
    "count": 10,
    "lastSeenTime": 1586903359622
  },
  "Weinrebe": {
    "count": 5,
    "lastSeenTime": 1586834116405
  },
  "Raum": {
    "count": 6,
    "lastSeenTime": 1586955238992
  },
  "Clown": {
    "count": 8,
    "lastSeenTime": 1586907400222
  },
  "Granate": {
    "count": 5,
    "lastSeenTime": 1586919664805,
    "difficulty": 0.6,
    "difficultyWeight": 5
  },
  "William Wallace": {
    "count": 5,
    "lastSeenTime": 1586914080489
  },
  "Rückenschmerzen": {
    "count": 9,
    "lastSeenTime": 1586920872823
  },
  "Tasse": {
    "count": 5,
    "lastSeenTime": 1586834116405,
    "difficulty": 0.25,
    "difficultyWeight": 8
  },
  "Nagelpflege": {
    "count": 6,
    "lastSeenTime": 1586914729019
  },
  "Torpedo": {
    "count": 5,
    "lastSeenTime": 1586956181125,
    "publicGameCount": 1,
    "difficulty": 0.5714285714285714,
    "difficultyWeight": 7
  },
  "Glitzer": {
    "count": 4,
    "lastSeenTime": 1586919460885,
    "difficulty": 0.8,
    "difficultyWeight": 10
  },
  "Wiesel": {
    "count": 9,
    "lastSeenTime": 1586955286110
  },
  "schweben": {
    "count": 9,
    "lastSeenTime": 1586919773316,
    "difficulty": 0.8378378378378378,
    "difficultyWeight": 37
  },
  "Mücke": {
    "count": 6,
    "lastSeenTime": 1586904248334
  },
  "Schatz": {
    "count": 6,
    "lastSeenTime": 1586815563232
  },
  "Hahn": {
    "count": 7,
    "lastSeenTime": 1586833201869,
    "difficulty": 0.75,
    "difficultyWeight": 32
  },
  "Essiggurke": {
    "count": 5,
    "lastSeenTime": 1586919610641
  },
  "Glühbirne": {
    "count": 4,
    "lastSeenTime": 1586913868342,
    "difficulty": 0.96,
    "difficultyWeight": 50
  },
  "Deodorant": {
    "count": 5,
    "lastSeenTime": 1586826254067,
    "difficulty": 0.4,
    "difficultyWeight": 5
  },
  "Harfe": {
    "count": 6,
    "lastSeenTime": 1586919327173,
    "difficulty": 0.8421052631578947,
    "difficultyWeight": 19
  },
  "umarmen": {
    "count": 7,
    "lastSeenTime": 1586921094172,
    "publicGameCount": 2,
    "difficulty": 0.9,
    "difficultyWeight": 10
  },
  "Diagonale": {
    "count": 6,
    "lastSeenTime": 1586900402067
  },
  "Axolotl": {
    "count": 5,
    "lastSeenTime": 1586827401622
  },
  "Radioaktivität": {
    "count": 9,
    "lastSeenTime": 1586958253918
  },
  "Kirby": {
    "count": 7,
    "lastSeenTime": 1586914387880
  },
  "Achterbahn": {
    "count": 5,
    "lastSeenTime": 1586914270411,
    "difficulty": 0.5333333333333333,
    "difficultyWeight": 15
  },
  "Hand": {
    "count": 6,
    "lastSeenTime": 1586900278764,
    "difficulty": 0.4,
    "difficultyWeight": 5
  },
  "Sattel": {
    "count": 8,
    "lastSeenTime": 1586958525123,
    "difficulty": 1,
    "difficultyWeight": 12
  },
  "Tacker": {
    "count": 4,
    "lastSeenTime": 1586815146910
  },
  "Vorstellungskraft": {
    "count": 3,
    "lastSeenTime": 1586959401878
  },
  "umkehren": {
    "count": 7,
    "lastSeenTime": 1586915974465
  },
  "Bürste": {
    "count": 6,
    "lastSeenTime": 1586954749821
  },
  "Regen": {
    "count": 2,
    "lastSeenTime": 1586569585076,
    "publicGameCount": 1,
    "difficulty": 0.8727272727272727,
    "difficultyWeight": 55
  },
  "aufwachen": {
    "count": 12,
    "lastSeenTime": 1586914869480,
    "difficulty": 1,
    "difficultyWeight": 9
  },
  "Kettensäge": {
    "count": 5,
    "lastSeenTime": 1586955634694
  },
  "Paris": {
    "count": 6,
    "lastSeenTime": 1586880936347
  },
  "Thunfisch": {
    "count": 8,
    "lastSeenTime": 1586902616627
  },
  "Picknick": {
    "count": 5,
    "lastSeenTime": 1586840420314,
    "difficulty": 1,
    "difficultyWeight": 12
  },
  "Nagelschere": {
    "count": 7,
    "lastSeenTime": 1586881247045
  },
  "Waschbrettbauch": {
    "count": 4,
    "lastSeenTime": 1586807932952
  },
  "Donald Duck": {
    "count": 6,
    "lastSeenTime": 1586903750797
  },
  "Hecht": {
    "count": 8,
    "lastSeenTime": 1586903048902
  },
  "Schallplatte": {
    "count": 4,
    "lastSeenTime": 1586826573783
  },
  "Laterne": {
    "count": 7,
    "lastSeenTime": 1586895872235,
    "difficulty": 1,
    "difficultyWeight": 9
  },
  "singen": {
    "count": 7,
    "lastSeenTime": 1586914181975
  },
  "Hampelmann": {
    "count": 7,
    "lastSeenTime": 1586805325864
  },
  "Krawatte": {
    "count": 3,
    "lastSeenTime": 1586815863706,
    "publicGameCount": 1,
    "difficulty": 1,
    "difficultyWeight": 8
  },
  "Werkzeugkasten": {
    "count": 9,
    "lastSeenTime": 1586919821927
  },
  "Qualle": {
    "count": 3,
    "lastSeenTime": 1586711756970,
    "difficulty": 0.8648648648648649,
    "difficultyWeight": 37,
    "publicGameCount": 1
  },
  "Gans": {
    "count": 6,
    "lastSeenTime": 1586897445244
  },
  "Flutlicht": {
    "count": 4,
    "lastSeenTime": 1586869078737,
    "difficulty": 1,
    "difficultyWeight": 8
  },
  "Dampf": {
    "count": 3,
    "lastSeenTime": 1586859475541,
    "difficulty": 0.25,
    "difficultyWeight": 8
  },
  "Reißverschlussverfahren": {
    "count": 4,
    "lastSeenTime": 1586913868342
  },
  "Polizist": {
    "count": 6,
    "lastSeenTime": 1586881085327,
    "difficulty": 1,
    "difficultyWeight": 6
  },
  "Aschenbecher": {
    "count": 7,
    "lastSeenTime": 1586879531452
  },
  "Drillinge": {
    "count": 4,
    "lastSeenTime": 1586897154134,
    "difficulty": 0.8275862068965517,
    "difficultyWeight": 29
  },
  "Insekt": {
    "count": 11,
    "lastSeenTime": 1586955938117
  },
  "Spongebob": {
    "count": 8,
    "lastSeenTime": 1586901003976
  },
  "Slalom": {
    "count": 9,
    "lastSeenTime": 1586903939910,
    "difficulty": 1,
    "difficultyWeight": 26
  },
  "Brieftaube": {
    "count": 7,
    "lastSeenTime": 1586919394915,
    "difficulty": 0.8571428571428571,
    "difficultyWeight": 7
  },
  "Gebäck": {
    "count": 6,
    "lastSeenTime": 1586879170331,
    "difficulty": 1,
    "difficultyWeight": 3
  },
  "Wasserhahn": {
    "count": 4,
    "lastSeenTime": 1586959207698
  },
  "tippen": {
    "count": 9,
    "lastSeenTime": 1586903673095
  },
  "Robin Hood": {
    "count": 4,
    "lastSeenTime": 1586873676367
  },
  "Lieferung": {
    "count": 5,
    "lastSeenTime": 1586806294106
  },
  "Strauß": {
    "count": 3,
    "lastSeenTime": 1586817693807
  },
  "fleischfressend": {
    "count": 4,
    "lastSeenTime": 1586816002310
  },
  "Faultier": {
    "count": 6,
    "lastSeenTime": 1586915856980
  },
  "Geburtstag": {
    "count": 6,
    "lastSeenTime": 1586806643897,
    "publicGameCount": 1,
    "difficulty": 0.9090909090909091,
    "difficultyWeight": 11
  },
  "Tierfutter": {
    "count": 4,
    "lastSeenTime": 1586901089526
  },
  "Obelix": {
    "count": 6,
    "lastSeenTime": 1586891242624,
    "difficulty": 0.2857142857142857,
    "difficultyWeight": 7
  },
  "Dollar": {
    "count": 10,
    "lastSeenTime": 1586955654939,
    "difficulty": 0.375,
    "difficultyWeight": 8
  },
  "Eichhörnchen": {
    "count": 6,
    "lastSeenTime": 1586958325841
  },
  "kratzen": {
    "count": 7,
    "lastSeenTime": 1586896509911,
    "difficulty": 1,
    "difficultyWeight": 5
  },
  "klatschen": {
    "count": 5,
    "lastSeenTime": 1586858550508
  },
  "Raumschiff": {
    "count": 3,
    "lastSeenTime": 1586713907372,
    "publicGameCount": 2,
    "difficulty": 1,
    "difficultyWeight": 7
  },
  "Ikea": {
    "count": 3,
    "lastSeenTime": 1586897039497
  },
  "sternhagelvoll": {
    "count": 5,
    "lastSeenTime": 1586812941735
  },
  "Keksdose": {
    "count": 14,
    "lastSeenTime": 1586955419506,
    "difficulty": 1,
    "difficultyWeight": 1
  },
  "Ameisenhaufen": {
    "count": 10,
    "lastSeenTime": 1586913424347
  },
  "Pilz": {
    "count": 5,
    "lastSeenTime": 1586957940022
  },
  "heiß": {
    "count": 5,
    "lastSeenTime": 1586868067356,
    "difficulty": 0.9761904761904762,
    "difficultyWeight": 42
  },
  "Tabak": {
    "count": 9,
    "lastSeenTime": 1586907386047,
    "difficulty": 1,
    "difficultyWeight": 10
  },
  "High Heels": {
    "count": 5,
    "lastSeenTime": 1586915139518
  },
  "Schublade": {
    "count": 3,
    "lastSeenTime": 1586874917114
  },
  "Coffeeshop": {
    "count": 5,
    "lastSeenTime": 1586896530203
  },
  "Schlauch": {
    "count": 5,
    "lastSeenTime": 1586889566614
  },
  "Teletubby": {
    "count": 7,
    "lastSeenTime": 1586902641439
  },
  "Butter": {
    "count": 10,
    "lastSeenTime": 1586889682224
  },
  "Münze": {
    "count": 3,
    "lastSeenTime": 1586909250467,
    "difficulty": 1,
    "difficultyWeight": 38
  },
  "Überschrift": {
    "count": 7,
    "lastSeenTime": 1586914883727,
    "difficulty": 0.5,
    "difficultyWeight": 6
  },
  "Rapunzel": {
    "count": 5,
    "lastSeenTime": 1586956130879
  },
  "Ofen": {
    "count": 10,
    "lastSeenTime": 1586908141001
  },
  "Afrofrisur": {
    "count": 5,
    "lastSeenTime": 1586913696698
  },
  "Aufkleber": {
    "count": 5,
    "lastSeenTime": 1586897455407
  },
  "magersüchtig": {
    "count": 7,
    "lastSeenTime": 1586919650407,
    "difficulty": 1,
    "difficultyWeight": 6
  },
  "Pflaume": {
    "count": 5,
    "lastSeenTime": 1586869723557
  },
  "Giftzwerg": {
    "count": 4,
    "lastSeenTime": 1586840557813,
    "difficulty": 0.9767441860465116,
    "difficultyWeight": 43
  },
  "Pfadfinder": {
    "count": 8,
    "lastSeenTime": 1586913606771
  },
  "Barcode": {
    "count": 9,
    "lastSeenTime": 1586956552197,
    "difficulty": 1,
    "difficultyWeight": 5
  },
  "Morgen": {
    "count": 7,
    "lastSeenTime": 1586818656558,
    "difficulty": 1,
    "difficultyWeight": 5
  },
  "Tisch": {
    "count": 9,
    "lastSeenTime": 1586955654939,
    "publicGameCount": 1,
    "difficulty": 0.42857142857142855,
    "difficultyWeight": 7
  },
  "Rotkehlchen": {
    "count": 11,
    "lastSeenTime": 1586914181975
  },
  "Kalender": {
    "count": 5,
    "lastSeenTime": 1586904040749
  },
  "Elektroauto": {
    "count": 6,
    "lastSeenTime": 1586955025294
  },
  "Jäger": {
    "count": 6,
    "lastSeenTime": 1586956019428,
    "difficulty": 1,
    "difficultyWeight": 7
  },
  "Hase": {
    "count": 7,
    "lastSeenTime": 1586903333845
  },
  "Kiste": {
    "count": 4,
    "lastSeenTime": 1586879520233
  },
  "Käsekuchen": {
    "count": 11,
    "lastSeenTime": 1586909201701
  },
  "Kreuzfahrt": {
    "count": 7,
    "lastSeenTime": 1586959570753,
    "publicGameCount": 1,
    "difficulty": 0.33333333333333337,
    "difficultyWeight": 9
  },
  "Corn Dog": {
    "count": 6,
    "lastSeenTime": 1586826796918
  },
  "Parade": {
    "count": 10,
    "lastSeenTime": 1586915442369,
    "publicGameCount": 1,
    "difficulty": 0.875,
    "difficultyWeight": 8
  },
  "Überfall": {
    "count": 9,
    "lastSeenTime": 1586919650407
  },
  "Quelle": {
    "count": 6,
    "lastSeenTime": 1586873482759
  },
  "Kontrabass": {
    "count": 11,
    "lastSeenTime": 1586919348003,
    "difficulty": 1,
    "difficultyWeight": 10
  },
  "Matratze": {
    "count": 8,
    "lastSeenTime": 1586920957237
  },
  "Xylofon": {
    "count": 8,
    "lastSeenTime": 1586955429782
  },
  "Daumen": {
    "count": 5,
    "lastSeenTime": 1586916404987,
    "publicGameCount": 1,
    "difficulty": 0.95,
    "difficultyWeight": 20
  },
  "Tischdecke": {
    "count": 9,
    "lastSeenTime": 1586958983867
  },
  "Dachschaden": {
    "count": 6,
    "lastSeenTime": 1586817605940
  },
  "Sonnenaufgang": {
    "count": 8,
    "lastSeenTime": 1586959789851,
    "publicGameCount": 2,
    "difficulty": 0.7777777777777778,
    "difficultyWeight": 27
  },
  "Halbleiter": {
    "count": 4,
    "lastSeenTime": 1586954678504
  },
  "Kranz": {
    "count": 8,
    "lastSeenTime": 1586908175585
  },
  "High Score": {
    "count": 8,
    "lastSeenTime": 1586874283356
  },
  "Fliegenklatsche": {
    "count": 9,
    "lastSeenTime": 1586920005337,
    "difficulty": 0.9622641509433962,
    "difficultyWeight": 53
  },
  "Maskottchen": {
    "count": 3,
    "lastSeenTime": 1586725969582,
    "difficulty": 0.9,
    "difficultyWeight": 10
  },
  "Tank": {
    "count": 10,
    "lastSeenTime": 1586913948292,
    "difficulty": 0.5555555555555556,
    "difficultyWeight": 9
  },
  "Kung Fu": {
    "count": 6,
    "lastSeenTime": 1586958325841
  },
  "lachen": {
    "count": 8,
    "lastSeenTime": 1586869397751,
    "difficulty": 0.25,
    "difficultyWeight": 8
  },
  "Klarinette": {
    "count": 5,
    "lastSeenTime": 1586896444515
  },
  "Hacker": {
    "count": 8,
    "lastSeenTime": 1586904412233
  },
  "vertikal": {
    "count": 5,
    "lastSeenTime": 1586959482662,
    "difficulty": 1,
    "difficultyWeight": 4
  },
  "Sicherheitsgurt": {
    "count": 6,
    "lastSeenTime": 1586903950072
  },
  "Schokolade": {
    "count": 9,
    "lastSeenTime": 1586913948292
  },
  "Schminke": {
    "count": 5,
    "lastSeenTime": 1586902926018
  },
  "Stoff": {
    "count": 5,
    "lastSeenTime": 1586858918703
  },
  "Grashüpfer": {
    "count": 6,
    "lastSeenTime": 1586916116567,
    "difficulty": 0.3333333333333333,
    "difficultyWeight": 12
  },
  "Schule": {
    "count": 7,
    "lastSeenTime": 1586914893907
  },
  "Schmied": {
    "count": 4,
    "lastSeenTime": 1586958077398
  },
  "erzählen": {
    "count": 6,
    "lastSeenTime": 1586920322622
  },
  "Weide": {
    "count": 8,
    "lastSeenTime": 1586901346838
  },
  "Golfwagen": {
    "count": 9,
    "lastSeenTime": 1586907900355
  },
  "Lichtschalter": {
    "count": 4,
    "lastSeenTime": 1586712768631
  },
  "Perle": {
    "count": 6,
    "lastSeenTime": 1586896769426,
    "difficulty": 0.9333333333333333,
    "difficultyWeight": 15,
    "publicGameCount": 1
  },
  "Explosion": {
    "count": 8,
    "lastSeenTime": 1586958873120,
    "publicGameCount": 1,
    "difficulty": 0.7333333333333333,
    "difficultyWeight": 15
  },
  "Kleid": {
    "count": 4,
    "lastSeenTime": 1586959417006,
    "difficulty": 0.8857142857142857,
    "difficultyWeight": 70
  },
  "Klingelton": {
    "count": 6,
    "lastSeenTime": 1586901606974
  },
  "depressiv": {
    "count": 6,
    "lastSeenTime": 1586958821795,
    "difficulty": 1,
    "difficultyWeight": 6
  },
  "Bauer": {
    "count": 5,
    "lastSeenTime": 1586827875892
  },
  "Bildschirm": {
    "count": 10,
    "lastSeenTime": 1586903836685
  },
  "Sparschwein": {
    "count": 15,
    "lastSeenTime": 1586914575777,
    "difficulty": 0.5714285714285714,
    "difficultyWeight": 7
  },
  "Lakritz": {
    "count": 9,
    "lastSeenTime": 1586914387880
  },
  "fahren": {
    "count": 8,
    "lastSeenTime": 1586890826443,
    "difficulty": 0.5,
    "difficultyWeight": 4
  },
  "Pistole": {
    "count": 4,
    "lastSeenTime": 1586907962530,
    "difficulty": 1,
    "difficultyWeight": 8
  },
  "Held": {
    "count": 6,
    "lastSeenTime": 1586907729754
  },
  "Eiffelturm": {
    "count": 4,
    "lastSeenTime": 1586874677647,
    "difficulty": 1,
    "difficultyWeight": 8
  },
  "Brücke": {
    "count": 10,
    "lastSeenTime": 1586908141001
  },
  "Kreis": {
    "count": 5,
    "lastSeenTime": 1586879705061
  },
  "Skalpell": {
    "count": 7,
    "lastSeenTime": 1586956215415
  },
  "Batterie": {
    "count": 10,
    "lastSeenTime": 1586954900699,
    "difficulty": 0.5333333333333333,
    "difficultyWeight": 15
  },
  "Hühnchen": {
    "count": 14,
    "lastSeenTime": 1586959738785
  },
  "Thron": {
    "count": 6,
    "lastSeenTime": 1586915365103,
    "difficulty": 1,
    "difficultyWeight": 5
  },
  "Narbe": {
    "count": 5,
    "lastSeenTime": 1586839235965
  },
  "Spielzeug": {
    "count": 8,
    "lastSeenTime": 1586915767461,
    "difficulty": 1,
    "difficultyWeight": 35
  },
  "Lunge": {
    "count": 5,
    "lastSeenTime": 1586815863706,
    "difficulty": 0.82,
    "difficultyWeight": 50
  },
  "Seepferdchen": {
    "count": 5,
    "lastSeenTime": 1586916074772
  },
  "melken": {
    "count": 4,
    "lastSeenTime": 1586874527694
  },
  "Paparazzi": {
    "count": 7,
    "lastSeenTime": 1586879892909
  },
  "Wange": {
    "count": 8,
    "lastSeenTime": 1586920312527,
    "difficulty": 0.30000000000000004,
    "difficultyWeight": 10
  },
  "Schneemann": {
    "count": 7,
    "lastSeenTime": 1586907351001,
    "difficulty": 0.5,
    "difficultyWeight": 14
  },
  "Muttermal": {
    "count": 5,
    "lastSeenTime": 1586908721366,
    "publicGameCount": 1,
    "difficulty": 0,
    "difficultyWeight": 1
  },
  "Groschen": {
    "count": 6,
    "lastSeenTime": 1586958052113
  },
  "differenzieren": {
    "count": 4,
    "lastSeenTime": 1586901181085
  },
  "Domino": {
    "count": 8,
    "lastSeenTime": 1586959482662,
    "difficulty": 0.4444444444444444,
    "difficultyWeight": 9
  },
  "schreien": {
    "count": 6,
    "lastSeenTime": 1586954588805
  },
  "Hawaiihemd": {
    "count": 4,
    "lastSeenTime": 1586874677647,
    "difficulty": 0.7222222222222222,
    "difficultyWeight": 18
  },
  "Brustkorb": {
    "count": 6,
    "lastSeenTime": 1586954967541,
    "difficulty": 1,
    "difficultyWeight": 7
  },
  "rechnen": {
    "count": 4,
    "lastSeenTime": 1586708948976,
    "difficulty": 0.8181818181818182,
    "difficultyWeight": 11
  },
  "Steckdose": {
    "count": 7,
    "lastSeenTime": 1586959178760,
    "difficulty": 0.75,
    "difficultyWeight": 20
  },
  "Gürtel": {
    "count": 2,
    "lastSeenTime": 1586874869836
  },
  "Hirsch": {
    "count": 8,
    "lastSeenTime": 1586956145689
  },
  "Pantomime": {
    "count": 6,
    "lastSeenTime": 1586858586343
  },
  "Rochen": {
    "count": 5,
    "lastSeenTime": 1586833743451,
    "difficulty": 0.7777777777777778,
    "difficultyWeight": 9
  },
  "Spiderman": {
    "count": 7,
    "lastSeenTime": 1586908697152,
    "difficulty": 0.8409090909090909,
    "difficultyWeight": 44
  },
  "Schlafenszeit": {
    "count": 7,
    "lastSeenTime": 1586916118321,
    "difficulty": 0.8,
    "difficultyWeight": 5
  },
  "Samsung": {
    "count": 10,
    "lastSeenTime": 1586896554510
  },
  "Nahrung": {
    "count": 5,
    "lastSeenTime": 1586957848254
  },
  "Harry Potter": {
    "count": 8,
    "lastSeenTime": 1586920722806,
    "difficulty": 1,
    "difficultyWeight": 24
  },
  "Katy Perry": {
    "count": 5,
    "lastSeenTime": 1586833719821
  },
  "Rückspiegel": {
    "count": 4,
    "lastSeenTime": 1586869324825
  },
  "Erdnuss-flips": {
    "count": 9,
    "lastSeenTime": 1586920690415
  },
  "Weltraum": {
    "count": 4,
    "lastSeenTime": 1586919697403,
    "difficulty": 1,
    "difficultyWeight": 12
  },
  "Virtual Reality": {
    "count": 7,
    "lastSeenTime": 1586869058242
  },
  "brüllen": {
    "count": 3,
    "lastSeenTime": 1586875005520
  },
  "Heinzelmännchen": {
    "count": 6,
    "lastSeenTime": 1586914560752,
    "difficulty": 1,
    "difficultyWeight": 8
  },
  "Spiel": {
    "count": 3,
    "lastSeenTime": 1586727314043
  },
  "Form": {
    "count": 3,
    "lastSeenTime": 1586921270600
  },
  "Küchentuch": {
    "count": 6,
    "lastSeenTime": 1586881021273
  },
  "Pilot": {
    "count": 6,
    "lastSeenTime": 1586958141882,
    "difficulty": 0.8545454545454545,
    "difficultyWeight": 55
  },
  "Gnom": {
    "count": 4,
    "lastSeenTime": 1586832821211
  },
  "Rettungsring": {
    "count": 6,
    "lastSeenTime": 1586914109258
  },
  "Augenbraue": {
    "count": 4,
    "lastSeenTime": 1586880956767,
    "difficulty": 1,
    "difficultyWeight": 3
  },
  "Schleichwerbung": {
    "count": 2,
    "lastSeenTime": 1586955457642
  },
  "Unfall": {
    "count": 6,
    "lastSeenTime": 1586920285874,
    "difficulty": 1,
    "difficultyWeight": 5
  },
  "Photosynthese": {
    "count": 3,
    "lastSeenTime": 1586955392320,
    "difficulty": 1,
    "difficultyWeight": 13
  },
  "Donald Trump": {
    "count": 6,
    "lastSeenTime": 1586901719438,
    "publicGameCount": 1,
    "difficulty": 1,
    "difficultyWeight": 4
  },
  "Farbpalette": {
    "count": 9,
    "lastSeenTime": 1586908721968
  },
  "Schlafstörung": {
    "count": 6,
    "lastSeenTime": 1586919858412
  },
  "Geldbeutel": {
    "count": 3,
    "lastSeenTime": 1586908494479
  },
  "Dreirad": {
    "count": 6,
    "lastSeenTime": 1586903421160
  },
  "Ziege": {
    "count": 6,
    "lastSeenTime": 1586826949191
  },
  "Palme": {
    "count": 4,
    "lastSeenTime": 1586908787449
  },
  "Grabstein": {
    "count": 7,
    "lastSeenTime": 1586867963918,
    "difficulty": 0.33333333333333337,
    "difficultyWeight": 9
  },
  "Wunde": {
    "count": 13,
    "lastSeenTime": 1586955152448
  },
  "Eule": {
    "count": 8,
    "lastSeenTime": 1586873598800
  },
  "Huf": {
    "count": 8,
    "lastSeenTime": 1586903867648
  },
  "Gourmet": {
    "count": 1,
    "lastSeenTime": 1586564694384
  },
  "Emoji": {
    "count": 3,
    "lastSeenTime": 1586633980478,
    "difficulty": 0.7941176470588235,
    "difficultyWeight": 34
  },
  "Stereo": {
    "count": 6,
    "lastSeenTime": 1586959595986,
    "difficulty": 0.9333333333333333,
    "difficultyWeight": 15
  },
  "Ski": {
    "count": 5,
    "lastSeenTime": 1586840301535
  },
  "Brezel": {
    "count": 5,
    "lastSeenTime": 1586921071907,
    "difficulty": 0.625,
    "difficultyWeight": 8
  },
  "Leine": {
    "count": 6,
    "lastSeenTime": 1586890378764
  },
  "Gepäckträger": {
    "count": 8,
    "lastSeenTime": 1586959279701,
    "difficulty": 1,
    "difficultyWeight": 36
  },
  "Jaguar": {
    "count": 8,
    "lastSeenTime": 1586920336794
  },
  "Paypal": {
    "count": 9,
    "lastSeenTime": 1586901166652,
    "difficulty": 1,
    "difficultyWeight": 22
  },
  "Spitzmaus": {
    "count": 5,
    "lastSeenTime": 1586921209345,
    "difficulty": 1,
    "difficultyWeight": 9
  },
  "Glas": {
    "count": 11,
    "lastSeenTime": 1586916328101,
    "difficulty": 1,
    "difficultyWeight": 1
  },
  "Wasserkreislauf": {
    "count": 5,
    "lastSeenTime": 1586880595683,
    "difficulty": 1,
    "difficultyWeight": 8
  },
  "Veröffentlichung": {
    "count": 7,
    "lastSeenTime": 1586916382544
  },
  "Henne": {
    "count": 8,
    "lastSeenTime": 1586840782545
  },
  "Barbar": {
    "count": 6,
    "lastSeenTime": 1586875120930
  },
  "Happy Meal": {
    "count": 9,
    "lastSeenTime": 1586919433115
  },
  "Kleber": {
    "count": 7,
    "lastSeenTime": 1586816382821
  },
  "schenken": {
    "count": 4,
    "lastSeenTime": 1586916731964,
    "publicGameCount": 1,
    "difficulty": 0.8333333333333334,
    "difficultyWeight": 24
  },
  "Portugal": {
    "count": 8,
    "lastSeenTime": 1586914309641,
    "difficulty": 0.42857142857142855,
    "difficultyWeight": 7
  },
  "Stirnband": {
    "count": 6,
    "lastSeenTime": 1586813412560,
    "difficulty": 0.8333333333333334,
    "difficultyWeight": 6
  },
  "Tannenzapfen": {
    "count": 9,
    "lastSeenTime": 1586958325841
  },
  "Uniform": {
    "count": 8,
    "lastSeenTime": 1586907547051,
    "publicGameCount": 1,
    "difficulty": 0.8,
    "difficultyWeight": 5
  },
  "Aubergine": {
    "count": 7,
    "lastSeenTime": 1586958642478,
    "difficulty": 0.8571428571428571,
    "difficultyWeight": 49
  },
  "Bruce Lee": {
    "count": 4,
    "lastSeenTime": 1586908921509
  },
  "Schacht": {
    "count": 4,
    "lastSeenTime": 1586839026220
  },
  "Youtuber": {
    "count": 8,
    "lastSeenTime": 1586959804291
  },
  "Leiche": {
    "count": 5,
    "lastSeenTime": 1586954999969,
    "difficulty": 0.9,
    "difficultyWeight": 10
  },
  "Mond": {
    "count": 2,
    "lastSeenTime": 1586806443319
  },
  "Pyramide": {
    "count": 10,
    "lastSeenTime": 1586955716758,
    "difficulty": 0.5714285714285714,
    "difficultyWeight": 14
  },
  "Tandem": {
    "count": 4,
    "lastSeenTime": 1586890378764,
    "difficulty": 0.9333333333333333,
    "difficultyWeight": 15
  },
  "Rucksack": {
    "count": 6,
    "lastSeenTime": 1586914557302
  },
  "Reisepass": {
    "count": 5,
    "lastSeenTime": 1586921104275
  },
  "Seite": {
    "count": 7,
    "lastSeenTime": 1586954664216,
    "difficulty": 1,
    "difficultyWeight": 10
  },
  "Abzug": {
    "count": 6,
    "lastSeenTime": 1586919836074
  },
  "Hosentasche": {
    "count": 10,
    "lastSeenTime": 1586879853156,
    "difficulty": 0.9411764705882353,
    "difficultyWeight": 17
  },
  "Rübe": {
    "count": 4,
    "lastSeenTime": 1586880531589
  },
  "Katzenklo": {
    "count": 5,
    "lastSeenTime": 1586959508307
  },
  "Zirkel": {
    "count": 2,
    "lastSeenTime": 1586902738828,
    "difficulty": 1,
    "difficultyWeight": 6
  },
  "Kohlrübe": {
    "count": 6,
    "lastSeenTime": 1586879359526,
    "difficulty": 1,
    "difficultyWeight": 5
  },
  "Jacht": {
    "count": 5,
    "lastSeenTime": 1586900866878,
    "difficulty": 1,
    "difficultyWeight": 23
  },
  "Serviette": {
    "count": 4,
    "lastSeenTime": 1586902776556
  },
  "hart": {
    "count": 3,
    "lastSeenTime": 1586874767406
  },
  "Karate": {
    "count": 11,
    "lastSeenTime": 1586902738828,
    "difficulty": 0.75,
    "difficultyWeight": 4
  },
  "Diamant": {
    "count": 9,
    "lastSeenTime": 1586958289143,
    "publicGameCount": 1,
    "difficulty": 0.918918918918919,
    "difficultyWeight": 37
  },
  "Lagerfeuer": {
    "count": 8,
    "lastSeenTime": 1586903853302
  },
  "Deadpool": {
    "count": 6,
    "lastSeenTime": 1586901547884
  },
  "Universum": {
    "count": 6,
    "lastSeenTime": 1586959090713
  },
  "Bär": {
    "count": 8,
    "lastSeenTime": 1586895747389,
    "difficulty": 0.8,
    "difficultyWeight": 15,
    "publicGameCount": 1
  },
  "Verkehrskontrolle": {
    "count": 8,
    "lastSeenTime": 1586959789851,
    "difficulty": 1,
    "difficultyWeight": 19
  },
  "Komet": {
    "count": 8,
    "lastSeenTime": 1586916382544
  },
  "übergeben": {
    "count": 7,
    "lastSeenTime": 1586900739157,
    "difficulty": 1,
    "difficultyWeight": 16
  },
  "Röntgenstrahlung": {
    "count": 5,
    "lastSeenTime": 1586916565022
  },
  "Polo": {
    "count": 7,
    "lastSeenTime": 1586919552513
  },
  "Gott": {
    "count": 3,
    "lastSeenTime": 1586707906596,
    "difficulty": 1,
    "difficultyWeight": 1
  },
  "Nasenlöcher": {
    "count": 4,
    "lastSeenTime": 1586920836680,
    "difficulty": 0.8571428571428571,
    "difficultyWeight": 7
  },
  "Berlin": {
    "count": 5,
    "lastSeenTime": 1586868603483,
    "difficulty": 0.7222222222222222,
    "difficultyWeight": 18
  },
  "küssen": {
    "count": 10,
    "lastSeenTime": 1586919783421
  },
  "Schleim": {
    "count": 9,
    "lastSeenTime": 1586890237676
  },
  "Flipper": {
    "count": 2,
    "lastSeenTime": 1586907839176,
    "difficulty": 0.8,
    "difficultyWeight": 5
  },
  "Schubkarre": {
    "count": 4,
    "lastSeenTime": 1586783730394
  },
  "Airbag": {
    "count": 3,
    "lastSeenTime": 1586827613351,
    "difficulty": 0.5714285714285714,
    "difficultyWeight": 7
  },
  "Schwerkraft": {
    "count": 3,
    "lastSeenTime": 1586826965958
  },
  "Gürteltier": {
    "count": 5,
    "lastSeenTime": 1586711447154
  },
  "Tunnel": {
    "count": 6,
    "lastSeenTime": 1586908054437,
    "difficulty": 0.8888888888888888,
    "difficultyWeight": 63
  },
  "Freiheitsstatue": {
    "count": 5,
    "lastSeenTime": 1586908392957
  },
  "Idee": {
    "count": 7,
    "lastSeenTime": 1586957888650,
    "difficulty": 0.25,
    "difficultyWeight": 8
  },
  "Black Friday": {
    "count": 5,
    "lastSeenTime": 1586958299322,
    "publicGameCount": 1,
    "difficulty": 0.6666666666666666,
    "difficultyWeight": 12
  },
  "Parkplatz": {
    "count": 4,
    "lastSeenTime": 1586920954248,
    "difficulty": 0.8,
    "difficultyWeight": 5
  },
  "Video": {
    "count": 9,
    "lastSeenTime": 1586908315438,
    "difficulty": 0.8,
    "difficultyWeight": 10
  },
  "Speck": {
    "count": 3,
    "lastSeenTime": 1586710915055
  },
  "Nebel": {
    "count": 4,
    "lastSeenTime": 1586954808666
  },
  "Queue": {
    "count": 4,
    "lastSeenTime": 1586859189711
  },
  "Stegosaurus": {
    "count": 6,
    "lastSeenTime": 1586916368116
  },
  "Hälfte": {
    "count": 4,
    "lastSeenTime": 1586920180037
  },
  "ernten": {
    "count": 7,
    "lastSeenTime": 1586873461546
  },
  "pieksen": {
    "count": 7,
    "lastSeenTime": 1586907414380
  },
  "Zahnarzt": {
    "count": 7,
    "lastSeenTime": 1586916153776,
    "difficulty": 0.7142857142857143,
    "difficultyWeight": 7
  },
  "Brownie": {
    "count": 6,
    "lastSeenTime": 1586914683667
  },
  "Busch": {
    "count": 5,
    "lastSeenTime": 1586879770813
  },
  "Fluss": {
    "count": 9,
    "lastSeenTime": 1586919384885
  },
  "Besen": {
    "count": 5,
    "lastSeenTime": 1586812283796
  },
  "Federball": {
    "count": 5,
    "lastSeenTime": 1586900264634,
    "difficulty": 0.8125,
    "difficultyWeight": 16
  },
  "Keyboard": {
    "count": 8,
    "lastSeenTime": 1586915609902,
    "difficulty": 0.7333333333333333,
    "difficultyWeight": 15
  },
  "James Bond": {
    "count": 6,
    "lastSeenTime": 1586904332915,
    "difficulty": 1,
    "difficultyWeight": 3
  },
  "Armleuchter": {
    "count": 6,
    "lastSeenTime": 1586919327173
  },
  "Litschi": {
    "count": 6,
    "lastSeenTime": 1586913581904,
    "publicGameCount": 1,
    "difficulty": 1,
    "difficultyWeight": 12
  },
  "Barbier": {
    "count": 9,
    "lastSeenTime": 1586902713347
  },
  "Sweatshirt": {
    "count": 7,
    "lastSeenTime": 1586907547051
  },
  "Fenster": {
    "count": 2,
    "lastSeenTime": 1586632836406,
    "difficulty": 0.4,
    "difficultyWeight": 10
  },
  "Lavendel": {
    "count": 3,
    "lastSeenTime": 1586889943537
  },
  "Taille": {
    "count": 7,
    "lastSeenTime": 1586908315438
  },
  "Befehlshaber": {
    "count": 5,
    "lastSeenTime": 1586904102272
  },
  "falsche Zähne": {
    "count": 6,
    "lastSeenTime": 1586915031903
  },
  "Sonic": {
    "count": 5,
    "lastSeenTime": 1586817982440
  },
  "Eber": {
    "count": 6,
    "lastSeenTime": 1586812232504
  },
  "Tonleiter": {
    "count": 11,
    "lastSeenTime": 1586908392957,
    "difficulty": 1,
    "difficultyWeight": 26
  },
  "Döner": {
    "count": 8,
    "lastSeenTime": 1586919797630,
    "difficulty": 1,
    "difficultyWeight": 3
  },
  "Dracula": {
    "count": 5,
    "lastSeenTime": 1586901152323,
    "difficulty": 1,
    "difficultyWeight": 4
  },
  "Seife": {
    "count": 5,
    "lastSeenTime": 1586913596578
  },
  "Feldstecher": {
    "count": 5,
    "lastSeenTime": 1586873383077
  },
  "Saft": {
    "count": 10,
    "lastSeenTime": 1586908697152,
    "difficulty": 1,
    "difficultyWeight": 12
  },
  "Portal": {
    "count": 7,
    "lastSeenTime": 1586880741697
  },
  "Badeanzug": {
    "count": 4,
    "lastSeenTime": 1586897368245,
    "publicGameCount": 1
  },
  "Trüffel": {
    "count": 6,
    "lastSeenTime": 1586919927137
  },
  "Bräutigam": {
    "count": 10,
    "lastSeenTime": 1586919836074,
    "difficulty": 1,
    "difficultyWeight": 26
  },
  "Baklava": {
    "count": 6,
    "lastSeenTime": 1586907693323
  },
  "studieren": {
    "count": 5,
    "lastSeenTime": 1586914206725,
    "difficulty": 0.9130434782608695,
    "difficultyWeight": 46
  },
  "Norwegen": {
    "count": 9,
    "lastSeenTime": 1586954588805
  },
  "Fidget Spinner": {
    "count": 7,
    "lastSeenTime": 1586901014103,
    "difficulty": 0.9090909090909091,
    "difficultyWeight": 11
  },
  "Schwanz": {
    "count": 9,
    "lastSeenTime": 1586915529401,
    "difficulty": 0.9285714285714286,
    "difficultyWeight": 14
  },
  "schnell": {
    "count": 10,
    "lastSeenTime": 1586909080709,
    "difficulty": 0.33333333333333337,
    "difficultyWeight": 9
  },
  "Grubenarbeiter": {
    "count": 3,
    "lastSeenTime": 1586838964600
  },
  "flach": {
    "count": 2,
    "lastSeenTime": 1586590610487
  },
  "Maki": {
    "count": 8,
    "lastSeenTime": 1586919433115
  },
  "Xbox": {
    "count": 5,
    "lastSeenTime": 1586954617410,
    "difficulty": 1,
    "difficultyWeight": 3
  },
  "Holz": {
    "count": 5,
    "lastSeenTime": 1586916731964
  },
  "Grippe": {
    "count": 4,
    "lastSeenTime": 1586913633876
  },
  "spannen": {
    "count": 6,
    "lastSeenTime": 1586890968538
  },
  "Allee": {
    "count": 5,
    "lastSeenTime": 1586868169552
  },
  "Anker": {
    "count": 3,
    "lastSeenTime": 1586890117788,
    "difficulty": 0.8571428571428571,
    "difficultyWeight": 21
  },
  "Stoppuhr": {
    "count": 7,
    "lastSeenTime": 1586958959180
  },
  "Gesundheit": {
    "count": 7,
    "lastSeenTime": 1586914820916,
    "publicGameCount": 1,
    "difficulty": 0.3333333333333333,
    "difficultyWeight": 6
  },
  "Tattoo": {
    "count": 3,
    "lastSeenTime": 1586874552382,
    "difficulty": 1,
    "difficultyWeight": 8
  },
  "Alkohol": {
    "count": 13,
    "lastSeenTime": 1586896554510,
    "difficulty": 0.33333333333333337,
    "difficultyWeight": 9
  },
  "Müller": {
    "count": 4,
    "lastSeenTime": 1586706924436
  },
  "Schweiz": {
    "count": 6,
    "lastSeenTime": 1586921032863
  },
  "verlassen": {
    "count": 5,
    "lastSeenTime": 1586840121442
  },
  "Pfannenwender": {
    "count": 11,
    "lastSeenTime": 1586879170331
  },
  "Kamel": {
    "count": 3,
    "lastSeenTime": 1586813676447,
    "difficulty": 0.9090909090909091,
    "difficultyWeight": 44
  },
  "Raupe": {
    "count": 8,
    "lastSeenTime": 1586818853521
  },
  "Hexe": {
    "count": 5,
    "lastSeenTime": 1586869324825,
    "difficulty": 1,
    "difficultyWeight": 11
  },
  "Vorschlaghammer": {
    "count": 4,
    "lastSeenTime": 1586920483192
  },
  "Dora": {
    "count": 5,
    "lastSeenTime": 1586959193290
  },
  "rollen": {
    "count": 8,
    "lastSeenTime": 1586896877511
  },
  "epilieren": {
    "count": 7,
    "lastSeenTime": 1586890675024
  },
  "herunterladen": {
    "count": 4,
    "lastSeenTime": 1586908155214,
    "difficulty": 1,
    "difficultyWeight": 3
  },
  "Lampe": {
    "count": 4,
    "lastSeenTime": 1586709093864,
    "difficulty": 0.6153846153846154,
    "difficultyWeight": 13
  },
  "Diskette": {
    "count": 7,
    "lastSeenTime": 1586955763543
  },
  "Umhang": {
    "count": 8,
    "lastSeenTime": 1586957874303
  },
  "Jeep": {
    "count": 6,
    "lastSeenTime": 1586954808666
  },
  "Horn": {
    "count": 15,
    "lastSeenTime": 1586915299932,
    "publicGameCount": 1,
    "difficulty": 0.7857142857142857,
    "difficultyWeight": 28
  },
  "Goldkette": {
    "count": 6,
    "lastSeenTime": 1586896854975,
    "difficulty": 1,
    "difficultyWeight": 11
  },
  "Kissen": {
    "count": 9,
    "lastSeenTime": 1586958794944
  },
  "Sechserpack": {
    "count": 6,
    "lastSeenTime": 1586921241738
  },
  "Keks": {
    "count": 5,
    "lastSeenTime": 1586956166526,
    "difficulty": 1,
    "difficultyWeight": 8
  },
  "Kettenkarussell": {
    "count": 5,
    "lastSeenTime": 1586913518787,
    "difficulty": 1,
    "difficultyWeight": 6
  },
  "Zeus": {
    "count": 10,
    "lastSeenTime": 1586919337853,
    "publicGameCount": 1,
    "difficulty": 1,
    "difficultyWeight": 24
  },
  "Leuchtstab": {
    "count": 8,
    "lastSeenTime": 1586919423610,
    "difficulty": 1,
    "difficultyWeight": 6
  },
  "Eisen": {
    "count": 7,
    "lastSeenTime": 1586955877308,
    "publicGameCount": 1,
    "difficulty": 0.5714285714285714,
    "difficultyWeight": 14
  },
  "Fee": {
    "count": 5,
    "lastSeenTime": 1586832498107
  },
  "Seilspringen": {
    "count": 4,
    "lastSeenTime": 1586805458421,
    "difficulty": 1,
    "difficultyWeight": 11
  },
  "Kopfschmerzen": {
    "count": 6,
    "lastSeenTime": 1586908757151
  },
  "Feuerwerk": {
    "count": 2,
    "lastSeenTime": 1586706242194,
    "difficulty": 0.6666666666666667,
    "difficultyWeight": 9
  },
  "Snoopy": {
    "count": 7,
    "lastSeenTime": 1586959401878,
    "difficulty": 1,
    "difficultyWeight": 6
  },
  "Waschbecken": {
    "count": 5,
    "lastSeenTime": 1586955923939,
    "difficulty": 0.75,
    "difficultyWeight": 4
  },
  "Karotte": {
    "count": 9,
    "lastSeenTime": 1586900941854,
    "difficulty": 1,
    "difficultyWeight": 11
  },
  "Dachboden": {
    "count": 5,
    "lastSeenTime": 1586875174443,
    "difficulty": 1,
    "difficultyWeight": 24
  },
  "Löffel": {
    "count": 16,
    "lastSeenTime": 1586959570753,
    "difficulty": 0.625,
    "difficultyWeight": 24
  },
  "Pups": {
    "count": 10,
    "lastSeenTime": 1586919522211,
    "difficulty": 1,
    "difficultyWeight": 3
  },
  "Stuhl": {
    "count": 5,
    "lastSeenTime": 1586956504597,
    "difficulty": 0.11111111111111116,
    "difficultyWeight": 9
  },
  "Vorhang": {
    "count": 8,
    "lastSeenTime": 1586955862348
  },
  "Abschleppwagen": {
    "count": 3,
    "lastSeenTime": 1586957989157
  },
  "Geier": {
    "count": 9,
    "lastSeenTime": 1586955209647
  },
  "rülpsen": {
    "count": 5,
    "lastSeenTime": 1586955138676,
    "difficulty": 0.8974358974358975,
    "difficultyWeight": 39
  },
  "Venusfliegenfalle": {
    "count": 6,
    "lastSeenTime": 1586958769937
  },
  "Klebeband": {
    "count": 9,
    "lastSeenTime": 1586891163944
  },
  "Tomate": {
    "count": 9,
    "lastSeenTime": 1586915529401
  },
  "Morse Code": {
    "count": 5,
    "lastSeenTime": 1586814103524,
    "difficulty": 1,
    "difficultyWeight": 16
  },
  "Morgan Freeman": {
    "count": 6,
    "lastSeenTime": 1586897139972,
    "difficulty": 0.14285714285714293,
    "difficultyWeight": 7
  },
  "sanft": {
    "count": 3,
    "lastSeenTime": 1586727828851
  },
  "Affe": {
    "count": 7,
    "lastSeenTime": 1586907839176
  },
  "Sieg": {
    "count": 1,
    "lastSeenTime": 1586567404446
  },
  "Feder": {
    "count": 8,
    "lastSeenTime": 1586901511433,
    "difficulty": 0.8260869565217391,
    "difficultyWeight": 23,
    "publicGameCount": 1
  },
  "niedlich": {
    "count": 5,
    "lastSeenTime": 1586860007509
  },
  "Bücherregal": {
    "count": 5,
    "lastSeenTime": 1586915417224,
    "difficulty": 1,
    "difficultyWeight": 8
  },
  "Frost": {
    "count": 6,
    "lastSeenTime": 1586959608297
  },
  "Melone": {
    "count": 4,
    "lastSeenTime": 1586828046616,
    "difficulty": 0.24999999999999997,
    "difficultyWeight": 12
  },
  "Gummiwürmchen": {
    "count": 4,
    "lastSeenTime": 1586834124544
  },
  "laden": {
    "count": 6,
    "lastSeenTime": 1586956253026
  },
  "Kokon": {
    "count": 3,
    "lastSeenTime": 1586651382688
  },
  "kalt": {
    "count": 5,
    "lastSeenTime": 1586958510664
  },
  "Talentshow": {
    "count": 6,
    "lastSeenTime": 1586916048294,
    "difficulty": 0.9722222222222222,
    "difficultyWeight": 36
  },
  "Asteroid": {
    "count": 7,
    "lastSeenTime": 1586913452717,
    "difficulty": 0.9565217391304348,
    "difficultyWeight": 23
  },
  "Spanien": {
    "count": 7,
    "lastSeenTime": 1586908646355
  },
  "Gitarre": {
    "count": 9,
    "lastSeenTime": 1586955532027,
    "difficulty": 0.4444444444444444,
    "difficultyWeight": 9
  },
  "Kohle": {
    "count": 6,
    "lastSeenTime": 1586817852869,
    "difficulty": 0.8461538461538461,
    "difficultyWeight": 13
  },
  "Pauke": {
    "count": 8,
    "lastSeenTime": 1586890654439,
    "difficulty": 1,
    "difficultyWeight": 11
  },
  "Festival": {
    "count": 5,
    "lastSeenTime": 1586920932924
  },
  "Orbit": {
    "count": 4,
    "lastSeenTime": 1586826458119,
    "difficulty": 1,
    "difficultyWeight": 2
  },
  "London": {
    "count": 5,
    "lastSeenTime": 1586901786675
  },
  "Wirbelsäule": {
    "count": 4,
    "lastSeenTime": 1586903543797
  },
  "Essen": {
    "count": 3,
    "lastSeenTime": 1586722986815
  },
  "Publikum": {
    "count": 8,
    "lastSeenTime": 1586900387786
  },
  "Fassade": {
    "count": 9,
    "lastSeenTime": 1586915731157,
    "difficulty": 0.7954545454545454,
    "difficultyWeight": 44
  },
  "Tiramisu": {
    "count": 5,
    "lastSeenTime": 1586955503431,
    "difficulty": 0.7777777777777778,
    "difficultyWeight": 9
  },
  "Eierschachtel": {
    "count": 4,
    "lastSeenTime": 1586955310334
  },
  "Schlagsahne": {
    "count": 5,
    "lastSeenTime": 1586959508307
  },
  "Laubsäge": {
    "count": 4,
    "lastSeenTime": 1586890704484,
    "difficulty": 1,
    "difficultyWeight": 7
  },
  "Discord": {
    "count": 7,
    "lastSeenTime": 1586958582243,
    "difficulty": 0.42857142857142855,
    "difficultyWeight": 14
  },
  "Bombe": {
    "count": 6,
    "lastSeenTime": 1586908304826,
    "difficulty": 0.6666666666666666,
    "difficultyWeight": 12
  },
  "Gold": {
    "count": 4,
    "lastSeenTime": 1586915503496
  },
  "Risiko": {
    "count": 5,
    "lastSeenTime": 1586955609964,
    "difficulty": 0.6,
    "difficultyWeight": 10
  },
  "Vampir": {
    "count": 5,
    "lastSeenTime": 1586958794944,
    "difficulty": 1,
    "difficultyWeight": 8
  },
  "McDonalds": {
    "count": 5,
    "lastSeenTime": 1586839555615
  },
  "Big Ben": {
    "count": 2,
    "lastSeenTime": 1586568977728
  },
  "Nussknacker": {
    "count": 7,
    "lastSeenTime": 1586860031219,
    "difficulty": 0.125,
    "difficultyWeight": 8
  },
  "Mutter": {
    "count": 5,
    "lastSeenTime": 1586909094851,
    "difficulty": 0.7931034482758621,
    "difficultyWeight": 29
  },
  "Huhn": {
    "count": 6,
    "lastSeenTime": 1586909201701,
    "difficulty": 0.9545454545454546,
    "difficultyWeight": 22
  },
  "Presslufthammer": {
    "count": 5,
    "lastSeenTime": 1586908368588
  },
  "Albtraum": {
    "count": 5,
    "lastSeenTime": 1586859971727,
    "difficulty": 0.6666666666666666,
    "difficultyWeight": 9
  },
  "Moskau": {
    "count": 7,
    "lastSeenTime": 1586903572367,
    "difficulty": 1,
    "difficultyWeight": 8
  },
  "Silo": {
    "count": 6,
    "lastSeenTime": 1586920397655,
    "difficulty": 1,
    "difficultyWeight": 10
  },
  "Otter": {
    "count": 6,
    "lastSeenTime": 1586816739369
  },
  "Untergrund": {
    "count": 8,
    "lastSeenTime": 1586900087077
  },
  "Antarktis": {
    "count": 10,
    "lastSeenTime": 1586955948209,
    "publicGameCount": 1,
    "difficulty": 0.875,
    "difficultyWeight": 8
  },
  "Königin": {
    "count": 4,
    "lastSeenTime": 1586880790147
  },
  "Brecheisen": {
    "count": 4,
    "lastSeenTime": 1586727600390,
    "difficulty": 0.6,
    "difficultyWeight": 10
  },
  "Rentner": {
    "count": 6,
    "lastSeenTime": 1586874242781
  },
  "Tretmühle": {
    "count": 4,
    "lastSeenTime": 1586723037365,
    "difficulty": 1,
    "difficultyWeight": 6
  },
  "Luftschiff": {
    "count": 7,
    "lastSeenTime": 1586858614527
  },
  "Gewitter": {
    "count": 6,
    "lastSeenTime": 1586907654631,
    "difficulty": 0.5714285714285714,
    "difficultyWeight": 7
  },
  "Fingerhut": {
    "count": 5,
    "lastSeenTime": 1586881031416
  },
  "Tierhandlung": {
    "count": 5,
    "lastSeenTime": 1586915503496
  },
  "Ägypten": {
    "count": 3,
    "lastSeenTime": 1586868093681
  },
  "schwitzen": {
    "count": 12,
    "lastSeenTime": 1586955010181
  },
  "John Cena": {
    "count": 9,
    "lastSeenTime": 1586900016201
  },
  "Türklinke": {
    "count": 7,
    "lastSeenTime": 1586909296939
  },
  "Garnele": {
    "count": 7,
    "lastSeenTime": 1586833778222,
    "difficulty": 1,
    "difficultyWeight": 3
  },
  "Gepard": {
    "count": 9,
    "lastSeenTime": 1586908165354
  },
  "Meteorit": {
    "count": 3,
    "lastSeenTime": 1586569956865
  },
  "stricken": {
    "count": 7,
    "lastSeenTime": 1586908541127
  },
  "predigen": {
    "count": 5,
    "lastSeenTime": 1586959767498
  },
  "Hello Kitty": {
    "count": 9,
    "lastSeenTime": 1586901067015,
    "difficulty": 0.33333333333333337,
    "difficultyWeight": 9
  },
  "Wunschliste": {
    "count": 11,
    "lastSeenTime": 1586891326336,
    "difficulty": 0.2857142857142857,
    "difficultyWeight": 7
  },
  "Stacheldraht": {
    "count": 6,
    "lastSeenTime": 1586890854897
  },
  "Fingerspitze": {
    "count": 8,
    "lastSeenTime": 1586899916722,
    "difficulty": 0.8333333333333334,
    "difficultyWeight": 6
  },
  "Orang-Utan": {
    "count": 5,
    "lastSeenTime": 1586873319498,
    "publicGameCount": 1,
    "difficulty": 0.45714285714285713,
    "difficultyWeight": 35
  },
  "Minion": {
    "count": 6,
    "lastSeenTime": 1586858650110,
    "difficulty": 1,
    "difficultyWeight": 44
  },
  "anzeigen": {
    "count": 6,
    "lastSeenTime": 1586901225801,
    "difficulty": 1,
    "difficultyWeight": 10
  },
  "Ballett": {
    "count": 3,
    "lastSeenTime": 1586817809043,
    "difficulty": 1,
    "difficultyWeight": 11
  },
  "Neuseeland": {
    "count": 5,
    "lastSeenTime": 1586921241738
  },
  "Schamesröte": {
    "count": 4,
    "lastSeenTime": 1586826573783
  },
  "Interview": {
    "count": 7,
    "lastSeenTime": 1586874880231,
    "difficulty": 0.7857142857142857,
    "difficultyWeight": 14
  },
  "Krankenschwester": {
    "count": 8,
    "lastSeenTime": 1586959218404,
    "difficulty": 1,
    "difficultyWeight": 2
  },
  "Kugelfisch": {
    "count": 4,
    "lastSeenTime": 1586900785753,
    "publicGameCount": 1,
    "difficulty": 0.8181818181818182,
    "difficultyWeight": 22
  },
  "Glück": {
    "count": 3,
    "lastSeenTime": 1586813998210
  },
  "Dexter": {
    "count": 11,
    "lastSeenTime": 1586913556716
  },
  "Schuh": {
    "count": 4,
    "lastSeenTime": 1586916116567
  },
  "gähnen": {
    "count": 7,
    "lastSeenTime": 1586919433115
  },
  "Socken": {
    "count": 8,
    "lastSeenTime": 1586914270411
  },
  "Dach": {
    "count": 6,
    "lastSeenTime": 1586914699787,
    "difficulty": 1,
    "difficultyWeight": 4
  },
  "Falte": {
    "count": 4,
    "lastSeenTime": 1586839830482
  },
  "Oscar": {
    "count": 5,
    "lastSeenTime": 1586914526456
  },
  "kauen": {
    "count": 7,
    "lastSeenTime": 1586897264698,
    "difficulty": 0.75,
    "difficultyWeight": 12
  },
  "Flughafen": {
    "count": 8,
    "lastSeenTime": 1586920446574,
    "difficulty": 1,
    "difficultyWeight": 4
  },
  "Chewbacca": {
    "count": 4,
    "lastSeenTime": 1586958471219
  },
  "Klimaanlage": {
    "count": 5,
    "lastSeenTime": 1586806917088
  },
  "Drama": {
    "count": 8,
    "lastSeenTime": 1586955238992,
    "difficulty": 0.8181818181818182,
    "difficultyWeight": 11
  },
  "Luigi": {
    "count": 7,
    "lastSeenTime": 1586920108914
  },
  "Teelicht": {
    "count": 7,
    "lastSeenTime": 1586920982151,
    "publicGameCount": 1
  },
  "Gießkanne": {
    "count": 3,
    "lastSeenTime": 1586958350172,
    "difficulty": 0.7666666666666667,
    "difficultyWeight": 60
  },
  "Spule": {
    "count": 8,
    "lastSeenTime": 1586913581904
  },
  "Papier": {
    "count": 6,
    "lastSeenTime": 1586900349607
  },
  "Kokosnuss": {
    "count": 4,
    "lastSeenTime": 1586859513029,
    "difficulty": 0.3333333333333333,
    "difficultyWeight": 6
  },
  "Cowboy": {
    "count": 4,
    "lastSeenTime": 1586904308260
  },
  "Community": {
    "count": 4,
    "lastSeenTime": 1586784190547
  },
  "Bodybuilding": {
    "count": 14,
    "lastSeenTime": 1586919954168,
    "difficulty": 1,
    "difficultyWeight": 7
  },
  "Vene": {
    "count": 3,
    "lastSeenTime": 1586718507669
  },
  "schreiben": {
    "count": 7,
    "lastSeenTime": 1586818118151,
    "difficulty": 0.375,
    "difficultyWeight": 8
  },
  "Medizin": {
    "count": 8,
    "lastSeenTime": 1586833955534
  },
  "Minecraft": {
    "count": 5,
    "lastSeenTime": 1586901346838
  },
  "gemütlich": {
    "count": 3,
    "lastSeenTime": 1586896791819,
    "difficulty": 0.7777777777777778,
    "difficultyWeight": 18
  },
  "Glockenspiel": {
    "count": 5,
    "lastSeenTime": 1586858799129,
    "difficulty": 0.5833333333333334,
    "difficultyWeight": 36
  },
  "Trommel": {
    "count": 4,
    "lastSeenTime": 1586921180981
  },
  "Notizblock": {
    "count": 1,
    "lastSeenTime": 1586568332043
  },
  "Knast": {
    "count": 5,
    "lastSeenTime": 1586916590016
  },
  "besuchen": {
    "count": 9,
    "lastSeenTime": 1586920722806,
    "publicGameCount": 1,
    "difficulty": 0.8333333333333334,
    "difficultyWeight": 6
  },
  "Kolibri": {
    "count": 5,
    "lastSeenTime": 1586914729019
  },
  "Amor": {
    "count": 7,
    "lastSeenTime": 1586839459663,
    "difficulty": 0.9523809523809523,
    "difficultyWeight": 21
  },
  "Windbeutel": {
    "count": 9,
    "lastSeenTime": 1586958100659,
    "difficulty": 0.875,
    "difficultyWeight": 8
  },
  "Joghurt": {
    "count": 8,
    "lastSeenTime": 1586915517726,
    "difficulty": 1,
    "difficultyWeight": 23
  },
  "sauber": {
    "count": 9,
    "lastSeenTime": 1586959714261,
    "difficulty": 0.8571428571428571,
    "difficultyWeight": 7
  },
  "Etagenbett": {
    "count": 6,
    "lastSeenTime": 1586915795632,
    "difficulty": 1,
    "difficultyWeight": 3
  },
  "Looping": {
    "count": 5,
    "lastSeenTime": 1586859083000
  },
  "Arzt": {
    "count": 4,
    "lastSeenTime": 1586958299322,
    "publicGameCount": 1,
    "difficulty": 0.5,
    "difficultyWeight": 4
  },
  "Handschuh": {
    "count": 8,
    "lastSeenTime": 1586901346838,
    "difficulty": 0.5555555555555556,
    "difficultyWeight": 9
  },
  "Tee": {
    "count": 8,
    "lastSeenTime": 1586896252833,
    "difficulty": 0.9375,
    "difficultyWeight": 64
  },
  "Gong": {
    "count": 6,
    "lastSeenTime": 1586897506792
  },
  "Jeans": {
    "count": 4,
    "lastSeenTime": 1586890675024,
    "difficulty": 0.9558823529411765,
    "difficultyWeight": 68
  },
  "Sonnenfinsternis": {
    "count": 6,
    "lastSeenTime": 1586921156587,
    "difficulty": 0.2,
    "difficultyWeight": 5
  },
  "Murmeltier": {
    "count": 4,
    "lastSeenTime": 1586807799590
  },
  "Pinsel": {
    "count": 9,
    "lastSeenTime": 1586958187063
  },
  "Angry Birds": {
    "count": 4,
    "lastSeenTime": 1586879770813
  },
  "Euter": {
    "count": 5,
    "lastSeenTime": 1586839296530,
    "difficulty": 0.7777777777777778,
    "difficultyWeight": 9
  },
  "Verband": {
    "count": 9,
    "lastSeenTime": 1586907522819,
    "difficulty": 1,
    "difficultyWeight": 7
  },
  "zusammenbrechen": {
    "count": 4,
    "lastSeenTime": 1586916116567
  },
  "Stein": {
    "count": 11,
    "lastSeenTime": 1586915846579
  },
  "Sekunde": {
    "count": 7,
    "lastSeenTime": 1586920883005,
    "difficulty": 0.6521739130434783,
    "difficultyWeight": 23
  },
  "Tupperdose": {
    "count": 6,
    "lastSeenTime": 1586895427458
  },
  "Hashtag": {
    "count": 5,
    "lastSeenTime": 1586920069685
  },
  "Fässchen": {
    "count": 5,
    "lastSeenTime": 1586920666083
  },
  "Roller": {
    "count": 3,
    "lastSeenTime": 1586956582710,
    "difficulty": 1,
    "difficultyWeight": 3
  },
  "zupfen": {
    "count": 4,
    "lastSeenTime": 1586874143406,
    "difficulty": 1,
    "difficultyWeight": 29
  },
  "Eidechse": {
    "count": 11,
    "lastSeenTime": 1586879622531
  },
  "Aufhänger": {
    "count": 6,
    "lastSeenTime": 1586920751201,
    "difficulty": 1,
    "difficultyWeight": 10
  },
  "Bohnenstange": {
    "count": 3,
    "lastSeenTime": 1586806443319,
    "difficulty": 1,
    "difficultyWeight": 12
  },
  "Eierbecher": {
    "count": 6,
    "lastSeenTime": 1586891038128
  },
  "Schiffstaufe": {
    "count": 3,
    "lastSeenTime": 1586879583717
  },
  "Pegasus": {
    "count": 6,
    "lastSeenTime": 1586955838030,
    "difficulty": 0.5555555555555556,
    "difficultyWeight": 9
  },
  "Kartoffelbrei": {
    "count": 10,
    "lastSeenTime": 1586897125781
  },
  "Nähmaschine": {
    "count": 8,
    "lastSeenTime": 1586900660239,
    "difficulty": 1,
    "difficultyWeight": 3
  },
  "Dusche": {
    "count": 9,
    "lastSeenTime": 1586908028334
  },
  "Zahnfee": {
    "count": 7,
    "lastSeenTime": 1586901812091,
    "difficulty": 0.8444444444444444,
    "difficultyWeight": 45
  },
  "Regenbogen": {
    "count": 9,
    "lastSeenTime": 1586954627819,
    "difficulty": 0.5714285714285714,
    "difficultyWeight": 7
  },
  "King Kong": {
    "count": 8,
    "lastSeenTime": 1586954798445
  },
  "Pause": {
    "count": 7,
    "lastSeenTime": 1586914716657,
    "difficulty": 0.5,
    "difficultyWeight": 8
  },
  "Helikopter": {
    "count": 12,
    "lastSeenTime": 1586956033779
  },
  "Windsack": {
    "count": 6,
    "lastSeenTime": 1586909226093
  },
  "Titanic": {
    "count": 5,
    "lastSeenTime": 1586901692884,
    "difficulty": 0.9130434782608695,
    "difficultyWeight": 23
  },
  "Gitter": {
    "count": 8,
    "lastSeenTime": 1586959031017
  },
  "Schinken": {
    "count": 3,
    "lastSeenTime": 1586873772670
  },
  "Yin Yang": {
    "count": 11,
    "lastSeenTime": 1586958859906,
    "publicGameCount": 1,
    "difficulty": 0.75,
    "difficultyWeight": 4
  },
  "Biene": {
    "count": 8,
    "lastSeenTime": 1586921094172,
    "publicGameCount": 1,
    "difficulty": 0.8333333333333334,
    "difficultyWeight": 6
  },
  "Mr. Bean": {
    "count": 3,
    "lastSeenTime": 1586718534191,
    "difficulty": 1,
    "difficultyWeight": 1
  },
  "Regal": {
    "count": 7,
    "lastSeenTime": 1586920397655
  },
  "Dämon": {
    "count": 8,
    "lastSeenTime": 1586958212167,
    "difficulty": 1,
    "difficultyWeight": 10
  },
  "Harpune": {
    "count": 3,
    "lastSeenTime": 1586879479650
  },
  "Schoß": {
    "count": 9,
    "lastSeenTime": 1586958897546
  },
  "Kaugummi": {
    "count": 6,
    "lastSeenTime": 1586903915052,
    "difficulty": 0.847457627118644,
    "difficultyWeight": 59
  },
  "Ninja": {
    "count": 10,
    "lastSeenTime": 1586914699787,
    "difficulty": 1,
    "difficultyWeight": 8
  },
  "Tentakel": {
    "count": 4,
    "lastSeenTime": 1586867988600
  },
  "Sonnenbrille": {
    "count": 4,
    "lastSeenTime": 1586897178546
  },
  "Abflussrohr": {
    "count": 3,
    "lastSeenTime": 1586651439627
  },
  "Reißverschluss": {
    "count": 7,
    "lastSeenTime": 1586955456078,
    "difficulty": 1,
    "difficultyWeight": 9
  },
  "Spaghettimonster": {
    "count": 8,
    "lastSeenTime": 1586908028334,
    "difficulty": 1,
    "difficultyWeight": 28,
    "publicGameCount": 1
  },
  "Koch": {
    "count": 12,
    "lastSeenTime": 1586919858412
  },
  "Kuckuck": {
    "count": 10,
    "lastSeenTime": 1586914792295,
    "difficulty": 1,
    "difficultyWeight": 10
  },
  "Rohr": {
    "count": 3,
    "lastSeenTime": 1586955409372,
    "difficulty": 0.5384615384615384,
    "difficultyWeight": 13
  },
  "Spargel": {
    "count": 5,
    "lastSeenTime": 1586956238875,
    "publicGameCount": 1,
    "difficulty": 1,
    "difficultyWeight": 5
  },
  "Bäcker": {
    "count": 8,
    "lastSeenTime": 1586919983033
  },
  "Skateboardfahrer": {
    "count": 9,
    "lastSeenTime": 1586959728501
  },
  "Nagellack": {
    "count": 3,
    "lastSeenTime": 1586908470256
  },
  "Nashorn": {
    "count": 4,
    "lastSeenTime": 1586826260649,
    "difficulty": 0.25,
    "difficultyWeight": 4
  },
  "Scooby Doo": {
    "count": 7,
    "lastSeenTime": 1586915806469,
    "publicGameCount": 1,
    "difficulty": 0.9642857142857143,
    "difficultyWeight": 28
  },
  "Dumbo": {
    "count": 3,
    "lastSeenTime": 1586633389254
  },
  "Tablette": {
    "count": 4,
    "lastSeenTime": 1586812594958,
    "difficulty": 0.75,
    "difficultyWeight": 16
  },
  "Brot": {
    "count": 8,
    "lastSeenTime": 1586959142159
  },
  "Brille": {
    "count": 6,
    "lastSeenTime": 1586817023533
  },
  "Taxi": {
    "count": 7,
    "lastSeenTime": 1586914080489,
    "difficulty": 1,
    "difficultyWeight": 8
  },
  "Croissant": {
    "count": 6,
    "lastSeenTime": 1586958748793,
    "difficulty": 1,
    "difficultyWeight": 30
  },
  "Zwillinge": {
    "count": 7,
    "lastSeenTime": 1586874360889,
    "publicGameCount": 1,
    "difficulty": 0.16666666666666663,
    "difficultyWeight": 6
  },
  "Minotaurus": {
    "count": 7,
    "lastSeenTime": 1586901274634,
    "difficulty": 1,
    "difficultyWeight": 10
  },
  "Supermarkt": {
    "count": 7,
    "lastSeenTime": 1586908816075,
    "difficulty": 1,
    "difficultyWeight": 5
  },
  "Schwert": {
    "count": 7,
    "lastSeenTime": 1586900311185
  },
  "Xerox": {
    "count": 8,
    "lastSeenTime": 1586840371821
  },
  "Möwe": {
    "count": 3,
    "lastSeenTime": 1586920407752
  },
  "Wolf": {
    "count": 7,
    "lastSeenTime": 1586957888650,
    "difficulty": 0.5714285714285714,
    "difficultyWeight": 7
  },
  "Lebkuchenhaus": {
    "count": 4,
    "lastSeenTime": 1586909046293
  },
  "Reinheit": {
    "count": 6,
    "lastSeenTime": 1586914742004
  },
  "Saftpresse": {
    "count": 3,
    "lastSeenTime": 1586718309085
  },
  "Heiße Schokolade": {
    "count": 3,
    "lastSeenTime": 1586900534749
  },
  "glücklich": {
    "count": 7,
    "lastSeenTime": 1586889682224,
    "difficulty": 1,
    "difficultyWeight": 32
  },
  "Drache": {
    "count": 8,
    "lastSeenTime": 1586955138676
  },
  "Rahmen": {
    "count": 5,
    "lastSeenTime": 1586859593426
  },
  "Malkasten": {
    "count": 6,
    "lastSeenTime": 1586920858569
  },
  "Seeigel": {
    "count": 5,
    "lastSeenTime": 1586879384676,
    "publicGameCount": 1,
    "difficulty": 0.30000000000000004,
    "difficultyWeight": 10
  },
  "drehen": {
    "count": 4,
    "lastSeenTime": 1586919312988
  },
  "Apfelkuchen": {
    "count": 6,
    "lastSeenTime": 1586813030635,
    "publicGameCount": 2,
    "difficulty": 0.8823529411764706,
    "difficultyWeight": 34
  },
  "Zwiebel": {
    "count": 5,
    "lastSeenTime": 1586915021163
  },
  "Gebiss": {
    "count": 8,
    "lastSeenTime": 1586955344691,
    "difficulty": 1,
    "difficultyWeight": 12
  },
  "Didgeridoo": {
    "count": 10,
    "lastSeenTime": 1586903543797
  },
  "gebrochenes Herz": {
    "count": 4,
    "lastSeenTime": 1586954975722,
    "difficulty": 1,
    "difficultyWeight": 2
  },
  "Falltür": {
    "count": 7,
    "lastSeenTime": 1586901521530,
    "difficulty": 0.8571428571428571,
    "difficultyWeight": 7
  },
  "Schlucht": {
    "count": 8,
    "lastSeenTime": 1586955503431,
    "difficulty": 0.7333333333333333,
    "difficultyWeight": 15
  },
  "Gewinner": {
    "count": 5,
    "lastSeenTime": 1586957855578
  },
  "Bestrafung": {
    "count": 5,
    "lastSeenTime": 1586902728303
  },
  "Dienstag": {
    "count": 6,
    "lastSeenTime": 1586806009929
  },
  "Vakuum": {
    "count": 6,
    "lastSeenTime": 1586959076238
  },
  "Wolkenkratzer": {
    "count": 3,
    "lastSeenTime": 1586859907934,
    "difficulty": 0.9090909090909091,
    "difficultyWeight": 22,
    "publicGameCount": 1
  },
  "Feierabend": {
    "count": 6,
    "lastSeenTime": 1586915985961
  },
  "Wirbel": {
    "count": 9,
    "lastSeenTime": 1586913768003
  },
  "Anhängerkupplung": {
    "count": 7,
    "lastSeenTime": 1586955899639
  },
  "Mikrowelle": {
    "count": 5,
    "lastSeenTime": 1586881081098,
    "publicGameCount": 1
  },
  "Kappe": {
    "count": 5,
    "lastSeenTime": 1586818561318,
    "difficulty": 0.7692307692307693,
    "difficultyWeight": 13
  },
  "Schweiß": {
    "count": 8,
    "lastSeenTime": 1586915406529
  },
  "Slinky": {
    "count": 3,
    "lastSeenTime": 1586869537106
  },
  "Bagel": {
    "count": 4,
    "lastSeenTime": 1586713516288
  },
  "Lehrer": {
    "count": 6,
    "lastSeenTime": 1586914918478
  },
  "Kreditkarte": {
    "count": 8,
    "lastSeenTime": 1586920872823
  },
  "Wetter": {
    "count": 6,
    "lastSeenTime": 1586920421902
  },
  "Polizei": {
    "count": 6,
    "lastSeenTime": 1586916575467,
    "difficulty": 1,
    "difficultyWeight": 11
  },
  "Pfau": {
    "count": 4,
    "lastSeenTime": 1586907788446
  },
  "Gras": {
    "count": 8,
    "lastSeenTime": 1586959344881
  },
  "Maßstab": {
    "count": 6,
    "lastSeenTime": 1586920576642
  },
  "Kontinent": {
    "count": 4,
    "lastSeenTime": 1586913816802
  },
  "Religion": {
    "count": 5,
    "lastSeenTime": 1586716941543
  },
  "Family Guy": {
    "count": 3,
    "lastSeenTime": 1586840180011
  },
  "Skydiving": {
    "count": 6,
    "lastSeenTime": 1586832445534
  },
  "Bergkette": {
    "count": 4,
    "lastSeenTime": 1586833456845
  },
  "Waschbär": {
    "count": 7,
    "lastSeenTime": 1586896816116
  },
  "Erbsen": {
    "count": 3,
    "lastSeenTime": 1586807947102
  },
  "Usain Bolt": {
    "count": 9,
    "lastSeenTime": 1586908597667
  },
  "Berg": {
    "count": 3,
    "lastSeenTime": 1586914683667
  },
  "Krebs": {
    "count": 5,
    "lastSeenTime": 1586957863554,
    "difficulty": 0.375,
    "difficultyWeight": 8
  },
  "Las Vegas": {
    "count": 5,
    "lastSeenTime": 1586860284075
  },
  "Finger": {
    "count": 4,
    "lastSeenTime": 1586903070812
  },
  "Pfeil": {
    "count": 7,
    "lastSeenTime": 1586869144958,
    "difficulty": 0.7727272727272727,
    "difficultyWeight": 22
  },
  "Gefrierschrank": {
    "count": 2,
    "lastSeenTime": 1586920226877
  },
  "Frühstück": {
    "count": 3,
    "lastSeenTime": 1586833915389
  },
  "Maske": {
    "count": 4,
    "lastSeenTime": 1586889552321,
    "difficulty": 0.85,
    "difficultyWeight": 20
  },
  "rasieren": {
    "count": 4,
    "lastSeenTime": 1586869412000,
    "publicGameCount": 1,
    "difficulty": 0.631578947368421,
    "difficultyWeight": 19
  },
  "Punktestand": {
    "count": 8,
    "lastSeenTime": 1586909030886
  },
  "Honig": {
    "count": 6,
    "lastSeenTime": 1586896569941,
    "difficulty": 0.5,
    "difficultyWeight": 8
  },
  "Haarspange": {
    "count": 4,
    "lastSeenTime": 1586874324510
  },
  "Sarg": {
    "count": 6,
    "lastSeenTime": 1586921071907,
    "difficulty": 0.375,
    "difficultyWeight": 8
  },
  "Der Rosarote Panther": {
    "count": 6,
    "lastSeenTime": 1586827465786
  },
  "Park": {
    "count": 2,
    "lastSeenTime": 1586880262858,
    "difficulty": 0.8181818181818182,
    "difficultyWeight": 11
  },
  "Rettungsweste": {
    "count": 6,
    "lastSeenTime": 1586907351001
  },
  "träumen": {
    "count": 4,
    "lastSeenTime": 1586873278128,
    "difficulty": 1,
    "difficultyWeight": 1
  },
  "Mopp": {
    "count": 4,
    "lastSeenTime": 1586827202213
  },
  "Jimmy Neutron": {
    "count": 7,
    "lastSeenTime": 1586920954248
  },
  "Wall-e": {
    "count": 7,
    "lastSeenTime": 1586903447898,
    "difficulty": 0.5,
    "difficultyWeight": 8
  },
  "Monster": {
    "count": 10,
    "lastSeenTime": 1586959131863,
    "difficulty": 0.7368421052631579,
    "difficultyWeight": 57
  },
  "Quadrat": {
    "count": 6,
    "lastSeenTime": 1586890615582,
    "difficulty": 0.782608695652174,
    "difficultyWeight": 23
  },
  "Limonade": {
    "count": 9,
    "lastSeenTime": 1586956166526
  },
  "Kaugummikugel": {
    "count": 10,
    "lastSeenTime": 1586913783574,
    "difficulty": 0.875,
    "difficultyWeight": 8
  },
  "reparieren": {
    "count": 7,
    "lastSeenTime": 1586914727219
  },
  "Adler": {
    "count": 9,
    "lastSeenTime": 1586954886360,
    "difficulty": 0.6666666666666666,
    "difficultyWeight": 3
  },
  "Aquarium": {
    "count": 7,
    "lastSeenTime": 1586909046293
  },
  "Teekanne": {
    "count": 9,
    "lastSeenTime": 1586958187063
  },
  "tanzen": {
    "count": 5,
    "lastSeenTime": 1586783155447
  },
  "Pinocchio": {
    "count": 4,
    "lastSeenTime": 1586955620355
  },
  "Hähnchen": {
    "count": 7,
    "lastSeenTime": 1586959608297
  },
  "Profi": {
    "count": 5,
    "lastSeenTime": 1586916469958
  },
  "Luftschloss": {
    "count": 4,
    "lastSeenTime": 1586717552913,
    "publicGameCount": 1,
    "difficulty": 0.8888888888888888,
    "difficultyWeight": 9
  },
  "Dänemark": {
    "count": 7,
    "lastSeenTime": 1586817683501
  },
  "Frucht": {
    "count": 8,
    "lastSeenTime": 1586915186679
  },
  "Gefängnis": {
    "count": 6,
    "lastSeenTime": 1586880581425
  },
  "Mammut": {
    "count": 3,
    "lastSeenTime": 1586919511658
  },
  "Durst": {
    "count": 3,
    "lastSeenTime": 1586880834242,
    "publicGameCount": 1,
    "difficulty": 0.875,
    "difficultyWeight": 8
  },
  "Lippenstift": {
    "count": 3,
    "lastSeenTime": 1586956336636
  },
  "Schweinchen Dick": {
    "count": 8,
    "lastSeenTime": 1586903726114
  },
  "Erdmännchen": {
    "count": 4,
    "lastSeenTime": 1586955848155
  },
  "Zelle": {
    "count": 7,
    "lastSeenTime": 1586915392140
  },
  "Krähe": {
    "count": 2,
    "lastSeenTime": 1586832774449,
    "publicGameCount": 1,
    "difficulty": 1,
    "difficultyWeight": 9
  },
  "Sonnenblume": {
    "count": 3,
    "lastSeenTime": 1586903736584,
    "difficulty": 0.4,
    "difficultyWeight": 5
  },
  "Hawaii": {
    "count": 8,
    "lastSeenTime": 1586956166526
  },
  "essen": {
    "count": 5,
    "lastSeenTime": 1586920957237
  },
  "Gefahr": {
    "count": 4,
    "lastSeenTime": 1586718488883
  },
  "Cyborg": {
    "count": 7,
    "lastSeenTime": 1586914600906,
    "difficulty": 1,
    "difficultyWeight": 4
  },
  "Italien": {
    "count": 2,
    "lastSeenTime": 1586707253394,
    "difficulty": 0.7000000000000001,
    "difficultyWeight": 10
  },
  "Türschloss": {
    "count": 6,
    "lastSeenTime": 1586916415277,
    "difficulty": 0.9285714285714286,
    "difficultyWeight": 28
  },
  "aufnehmen": {
    "count": 6,
    "lastSeenTime": 1586901486180
  },
  "Yoshi": {
    "count": 6,
    "lastSeenTime": 1586899955403
  },
  "Leiter": {
    "count": 11,
    "lastSeenTime": 1586958525123
  },
  "Vollkornbrot": {
    "count": 5,
    "lastSeenTime": 1586902983849
  },
  "Arbeitszimmer": {
    "count": 5,
    "lastSeenTime": 1586904322754,
    "difficulty": 0.9,
    "difficultyWeight": 10
  },
  "Turban": {
    "count": 6,
    "lastSeenTime": 1586873614695
  },
  "Landschaft": {
    "count": 5,
    "lastSeenTime": 1586958897546
  },
  "Rührei": {
    "count": 4,
    "lastSeenTime": 1586956552197,
    "difficulty": 0.8,
    "difficultyWeight": 5
  },
  "Steam": {
    "count": 6,
    "lastSeenTime": 1586916328101
  },
  "schlagen": {
    "count": 8,
    "lastSeenTime": 1586909296939
  },
  "Gleichgewicht": {
    "count": 9,
    "lastSeenTime": 1586955862348
  },
  "Bratsche": {
    "count": 5,
    "lastSeenTime": 1586868707935
  },
  "Herkules": {
    "count": 8,
    "lastSeenTime": 1586959850831,
    "difficulty": 1,
    "difficultyWeight": 8
  },
  "überwintern": {
    "count": 7,
    "lastSeenTime": 1586873642917
  },
  "Anubis": {
    "count": 6,
    "lastSeenTime": 1586902587809
  },
  "Känguru": {
    "count": 9,
    "lastSeenTime": 1586914904157,
    "difficulty": 1,
    "difficultyWeight": 53
  },
  "Erde": {
    "count": 6,
    "lastSeenTime": 1586916690790,
    "publicGameCount": 1,
    "difficulty": 0.30000000000000004,
    "difficultyWeight": 10
  },
  "Tiegel": {
    "count": 5,
    "lastSeenTime": 1586868982986
  },
  "Steinzeit": {
    "count": 8,
    "lastSeenTime": 1586956489403
  },
  "Brett": {
    "count": 8,
    "lastSeenTime": 1586959031017
  },
  "Puder": {
    "count": 7,
    "lastSeenTime": 1586859727901
  },
  "Räucherstäbchen": {
    "count": 5,
    "lastSeenTime": 1586959218404
  },
  "ausruhen": {
    "count": 3,
    "lastSeenTime": 1586782238298,
    "difficulty": 0.8,
    "difficultyWeight": 10
  },
  "Anwalt": {
    "count": 4,
    "lastSeenTime": 1586904135139,
    "difficulty": 1,
    "difficultyWeight": 3
  },
  "Indien": {
    "count": 8,
    "lastSeenTime": 1586814402764
  },
  "Wasserfarbkasten": {
    "count": 6,
    "lastSeenTime": 1586955948209
  },
  "Finnland": {
    "count": 5,
    "lastSeenTime": 1586955634694
  },
  "Susan Wojcicki": {
    "count": 10,
    "lastSeenTime": 1586921043126
  },
  "Handtuch": {
    "count": 7,
    "lastSeenTime": 1586957940022,
    "difficulty": 0.4444444444444444,
    "difficultyWeight": 9
  },
  "Antilope": {
    "count": 6,
    "lastSeenTime": 1586916415277
  },
  "Schleifpapier": {
    "count": 4,
    "lastSeenTime": 1586899930998
  },
  "Bart Simpson": {
    "count": 6,
    "lastSeenTime": 1586901606974
  },
  "Thaddäus Tentakel": {
    "count": 4,
    "lastSeenTime": 1586958848321,
    "difficulty": 0.7142857142857143,
    "difficultyWeight": 7
  },
  "riechen": {
    "count": 7,
    "lastSeenTime": 1586903587379
  },
  "Cupcake": {
    "count": 3,
    "lastSeenTime": 1586839889181
  },
  "Verlies": {
    "count": 6,
    "lastSeenTime": 1586858765870
  },
  "Maurer": {
    "count": 11,
    "lastSeenTime": 1586955429782,
    "difficulty": 1,
    "difficultyWeight": 7
  },
  "Phineas und Ferb": {
    "count": 11,
    "lastSeenTime": 1586880631457
  },
  "Zoo": {
    "count": 3,
    "lastSeenTime": 1586833481600
  },
  "Schiedsrichter": {
    "count": 8,
    "lastSeenTime": 1586889725951
  },
  "Vlogger": {
    "count": 4,
    "lastSeenTime": 1586833392800
  },
  "erröten": {
    "count": 5,
    "lastSeenTime": 1586816678295,
    "difficulty": 1,
    "difficultyWeight": 1
  },
  "John Lennon": {
    "count": 5,
    "lastSeenTime": 1586833992491
  },
  "Protest": {
    "count": 5,
    "lastSeenTime": 1586907461933
  },
  "Pedal": {
    "count": 6,
    "lastSeenTime": 1586896644364
  },
  "Stadion": {
    "count": 8,
    "lastSeenTime": 1586955098414,
    "difficulty": 0.75,
    "difficultyWeight": 8
  },
  "Pfad": {
    "count": 9,
    "lastSeenTime": 1586919327173
  },
  "Nagelfeile": {
    "count": 9,
    "lastSeenTime": 1586916430600
  },
  "Kleiderbügel": {
    "count": 8,
    "lastSeenTime": 1586959391653,
    "difficulty": 0.45454545454545453,
    "difficultyWeight": 11
  },
  "Kleinbus": {
    "count": 5,
    "lastSeenTime": 1586879867638
  },
  "Anhalter": {
    "count": 8,
    "lastSeenTime": 1586869459522
  },
  "Puzzle": {
    "count": 8,
    "lastSeenTime": 1586955595532
  },
  "Fächer": {
    "count": 4,
    "lastSeenTime": 1586895604165
  },
  "Bienenstock": {
    "count": 10,
    "lastSeenTime": 1586921298152
  },
  "Türstopper": {
    "count": 9,
    "lastSeenTime": 1586959508307,
    "difficulty": 1,
    "difficultyWeight": 1
  },
  "Honigwabe": {
    "count": 8,
    "lastSeenTime": 1586900549029
  },
  "dünn": {
    "count": 3,
    "lastSeenTime": 1586609722124,
    "difficulty": 0.875,
    "difficultyWeight": 24
  },
  "Japan": {
    "count": 5,
    "lastSeenTime": 1586839151208
  },
  "Mitfahrgelegenheit": {
    "count": 6,
    "lastSeenTime": 1586891149686
  },
  "Trauben": {
    "count": 5,
    "lastSeenTime": 1586879608242
  },
  "Schwertfisch": {
    "count": 4,
    "lastSeenTime": 1586915299932,
    "difficulty": 1,
    "difficultyWeight": 7
  },
  "Osterhase": {
    "count": 11,
    "lastSeenTime": 1586916164227,
    "difficulty": 1,
    "difficultyWeight": 15
  },
  "Elster": {
    "count": 3,
    "lastSeenTime": 1586901572201
  },
  "Bart": {
    "count": 4,
    "lastSeenTime": 1586901672587,
    "difficulty": 0.6363636363636364,
    "difficultyWeight": 11
  },
  "Bibliothekar": {
    "count": 10,
    "lastSeenTime": 1586955691477
  },
  "Olivenöl": {
    "count": 5,
    "lastSeenTime": 1586868185306
  },
  "Krug": {
    "count": 5,
    "lastSeenTime": 1586913893763
  },
  "Motte": {
    "count": 4,
    "lastSeenTime": 1586915129381,
    "difficulty": 1,
    "difficultyWeight": 4
  },
  "Meister": {
    "count": 6,
    "lastSeenTime": 1586900650119
  },
  "Südafrika": {
    "count": 8,
    "lastSeenTime": 1586919687264
  },
  "Pfund": {
    "count": 7,
    "lastSeenTime": 1586915985961
  },
  "Schimmel": {
    "count": 7,
    "lastSeenTime": 1586903359622,
    "publicGameCount": 1,
    "difficulty": 0.8181818181818182,
    "difficultyWeight": 11
  },
  "Decke": {
    "count": 5,
    "lastSeenTime": 1586914465537
  },
  "Superkraft": {
    "count": 4,
    "lastSeenTime": 1586874652224,
    "difficulty": 0.8571428571428571,
    "difficultyWeight": 14
  },
  "Blaubeere": {
    "count": 7,
    "lastSeenTime": 1586914515993
  },
  "Nerd": {
    "count": 4,
    "lastSeenTime": 1586890454038
  },
  "Pfeffersalami": {
    "count": 6,
    "lastSeenTime": 1586873383077
  },
  "Mount Rushmore": {
    "count": 6,
    "lastSeenTime": 1586890816340,
    "difficulty": 1,
    "difficultyWeight": 4
  },
  "Kamera": {
    "count": 7,
    "lastSeenTime": 1586915566245
  },
  "Reisender": {
    "count": 3,
    "lastSeenTime": 1586874966390
  },
  "Manege": {
    "count": 7,
    "lastSeenTime": 1586896030801
  },
  "Linse": {
    "count": 5,
    "lastSeenTime": 1586880201224
  },
  "Spieler": {
    "count": 7,
    "lastSeenTime": 1586901500696
  },
  "übersetzen": {
    "count": 6,
    "lastSeenTime": 1586890714906,
    "difficulty": 1,
    "difficultyWeight": 6
  },
  "Verkehr": {
    "count": 4,
    "lastSeenTime": 1586818927139
  },
  "Lachs": {
    "count": 7,
    "lastSeenTime": 1586958496125,
    "difficulty": 1,
    "difficultyWeight": 4
  },
  "Truhe": {
    "count": 6,
    "lastSeenTime": 1586858897962,
    "difficulty": 0.625,
    "difficultyWeight": 16
  },
  "Zuckerwatte": {
    "count": 3,
    "lastSeenTime": 1586958658325
  },
  "Oberfläche": {
    "count": 6,
    "lastSeenTime": 1586914893907
  },
  "William Shakespeare": {
    "count": 4,
    "lastSeenTime": 1586904147395
  },
  "Zitteraal": {
    "count": 7,
    "lastSeenTime": 1586839545523,
    "publicGameCount": 1,
    "difficulty": 0.5555555555555556,
    "difficultyWeight": 9
  },
  "Schweinestall": {
    "count": 4,
    "lastSeenTime": 1586959753084,
    "difficulty": 1,
    "difficultyWeight": 4
  },
  "Mars": {
    "count": 5,
    "lastSeenTime": 1586908982317,
    "publicGameCount": 1,
    "difficulty": 0,
    "difficultyWeight": 8
  },
  "Kürbis": {
    "count": 3,
    "lastSeenTime": 1586907386047
  },
  "Fahrbahn": {
    "count": 6,
    "lastSeenTime": 1586956477794
  },
  "Knie": {
    "count": 7,
    "lastSeenTime": 1586956016971
  },
  "Stiefel": {
    "count": 9,
    "lastSeenTime": 1586956181125,
    "difficulty": 0.9047619047619048,
    "difficultyWeight": 21
  },
  "Hunger": {
    "count": 13,
    "lastSeenTime": 1586957925802,
    "difficulty": 1,
    "difficultyWeight": 60
  },
  "Konsole": {
    "count": 3,
    "lastSeenTime": 1586724124838,
    "difficulty": 1,
    "difficultyWeight": 6
  },
  "Magier": {
    "count": 13,
    "lastSeenTime": 1586889696772
  },
  "Badekappe": {
    "count": 3,
    "lastSeenTime": 1586826659168
  },
  "Sommer": {
    "count": 6,
    "lastSeenTime": 1586874622713,
    "difficulty": 1,
    "difficultyWeight": 4
  },
  "abheben": {
    "count": 4,
    "lastSeenTime": 1586903771458
  },
  "Weintrauben": {
    "count": 8,
    "lastSeenTime": 1586908721968,
    "publicGameCount": 1,
    "difficulty": 0,
    "difficultyWeight": 8
  },
  "Band": {
    "count": 5,
    "lastSeenTime": 1586916352907
  },
  "Tablett": {
    "count": 10,
    "lastSeenTime": 1586907743895
  },
  "Fahrwerk": {
    "count": 4,
    "lastSeenTime": 1586908470256
  },
  "Dattel": {
    "count": 6,
    "lastSeenTime": 1586916493492
  },
  "Darwin": {
    "count": 4,
    "lastSeenTime": 1586833367687
  },
  "Festung": {
    "count": 2,
    "lastSeenTime": 1586914221389
  },
  "Haltestelle": {
    "count": 5,
    "lastSeenTime": 1586908763208,
    "difficulty": 1,
    "difficultyWeight": 6
  },
  "Kaktus": {
    "count": 6,
    "lastSeenTime": 1586889576969,
    "difficulty": 0.45454545454545453,
    "difficultyWeight": 11
  },
  "Nachos": {
    "count": 6,
    "lastSeenTime": 1586959714261,
    "publicGameCount": 1,
    "difficulty": 0.9090909090909091,
    "difficultyWeight": 11
  },
  "Stoßstange": {
    "count": 6,
    "lastSeenTime": 1586955848155
  },
  "Wattestäbchen": {
    "count": 3,
    "lastSeenTime": 1586913962978,
    "difficulty": 0.9111111111111111,
    "difficultyWeight": 45
  },
  "Ratatouille": {
    "count": 4,
    "lastSeenTime": 1586880239561
  },
  "Wonder Woman": {
    "count": 7,
    "lastSeenTime": 1586920947128
  },
  "Wanze": {
    "count": 6,
    "lastSeenTime": 1586920566543
  },
  "Vater": {
    "count": 7,
    "lastSeenTime": 1586890353925,
    "difficulty": 1,
    "difficultyWeight": 7
  },
  "Pappe": {
    "count": 11,
    "lastSeenTime": 1586919858412
  },
  "Augenbinde": {
    "count": 8,
    "lastSeenTime": 1586956263288
  },
  "Bank": {
    "count": 4,
    "lastSeenTime": 1586833496863
  },
  "Blume": {
    "count": 6,
    "lastSeenTime": 1586959547988,
    "difficulty": 1,
    "difficultyWeight": 9
  },
  "Mercedes": {
    "count": 5,
    "lastSeenTime": 1586880609937,
    "publicGameCount": 1,
    "difficulty": 0.9,
    "difficultyWeight": 10
  },
  "gehen": {
    "count": 6,
    "lastSeenTime": 1586895808940
  },
  "Angelina Jolie": {
    "count": 6,
    "lastSeenTime": 1586919711626
  },
  "Versicherung": {
    "count": 5,
    "lastSeenTime": 1586955877308
  },
  "Straße": {
    "count": 12,
    "lastSeenTime": 1586920627218
  },
  "USB": {
    "count": 8,
    "lastSeenTime": 1586904234099
  },
  "Aristokrat": {
    "count": 2,
    "lastSeenTime": 1586901089526
  },
  "Fischer": {
    "count": 5,
    "lastSeenTime": 1586959121500,
    "difficulty": 1,
    "difficultyWeight": 10
  },
  "Pavian": {
    "count": 6,
    "lastSeenTime": 1586907743895
  },
  "Wunderland": {
    "count": 4,
    "lastSeenTime": 1586958077398,
    "difficulty": 1,
    "difficultyWeight": 1
  },
  "Glut": {
    "count": 5,
    "lastSeenTime": 1586915215166
  },
  "Grenze": {
    "count": 8,
    "lastSeenTime": 1586907436625
  },
  "Gabel": {
    "count": 9,
    "lastSeenTime": 1586839654737,
    "difficulty": 0.4,
    "difficultyWeight": 15
  },
  "Cat Woman": {
    "count": 2,
    "lastSeenTime": 1586907707534
  },
  "Skyline": {
    "count": 5,
    "lastSeenTime": 1586954603172,
    "difficulty": 1,
    "difficultyWeight": 5
  },
  "Erdbeere": {
    "count": 7,
    "lastSeenTime": 1586915186679,
    "publicGameCount": 1,
    "difficulty": 0.7352941176470589,
    "difficultyWeight": 34
  },
  "Finn und Jake": {
    "count": 4,
    "lastSeenTime": 1586958572090
  },
  "Mel Gibson": {
    "count": 7,
    "lastSeenTime": 1586958658325
  },
  "Hausnummer": {
    "count": 7,
    "lastSeenTime": 1586908204031,
    "difficulty": 0.5,
    "difficultyWeight": 6
  },
  "Spitzhacke": {
    "count": 6,
    "lastSeenTime": 1586956019428
  },
  "Bajonett": {
    "count": 9,
    "lastSeenTime": 1586958141882
  },
  "Sumo": {
    "count": 5,
    "lastSeenTime": 1586895467027
  },
  "gelangweilt": {
    "count": 6,
    "lastSeenTime": 1586920947128
  },
  "Ziegenbart": {
    "count": 4,
    "lastSeenTime": 1586896444515,
    "difficulty": 0.8333333333333334,
    "difficultyWeight": 24
  },
  "Tetra Pak": {
    "count": 9,
    "lastSeenTime": 1586915821415,
    "difficulty": 1,
    "difficultyWeight": 7
  },
  "ringen": {
    "count": 5,
    "lastSeenTime": 1586813390059
  },
  "Salzwasser": {
    "count": 4,
    "lastSeenTime": 1586783782593,
    "difficulty": 1,
    "difficultyWeight": 7
  },
  "Zeitung": {
    "count": 2,
    "lastSeenTime": 1586916230527,
    "difficulty": 0.6,
    "difficultyWeight": 10
  },
  "Rad": {
    "count": 3,
    "lastSeenTime": 1586840557813,
    "difficulty": 0.9545454545454546,
    "difficultyWeight": 44
  },
  "Scheidung": {
    "count": 3,
    "lastSeenTime": 1586958755749
  },
  "Schneesturm": {
    "count": 7,
    "lastSeenTime": 1586955088291,
    "difficulty": 0.8571428571428571,
    "difficultyWeight": 7
  },
  "Lexikon": {
    "count": 4,
    "lastSeenTime": 1586897164255
  },
  "Haken": {
    "count": 2,
    "lastSeenTime": 1586784919921
  },
  "Mechaniker": {
    "count": 6,
    "lastSeenTime": 1586896926779
  },
  "Skulptur": {
    "count": 6,
    "lastSeenTime": 1586880423720,
    "difficulty": 1,
    "difficultyWeight": 8
  },
  "Server": {
    "count": 4,
    "lastSeenTime": 1586726377611
  },
  "Bademantel": {
    "count": 7,
    "lastSeenTime": 1586895639777,
    "publicGameCount": 1,
    "difficulty": 0.8709677419354839,
    "difficultyWeight": 31
  },
  "Griechenland": {
    "count": 3,
    "lastSeenTime": 1586902926018,
    "publicGameCount": 1,
    "difficulty": 0.9285714285714286,
    "difficultyWeight": 14
  },
  "trinken": {
    "count": 7,
    "lastSeenTime": 1586914631184
  },
  "schlecht": {
    "count": 6,
    "lastSeenTime": 1586896595013,
    "publicGameCount": 1
  },
  "vergessen": {
    "count": 7,
    "lastSeenTime": 1586955559002
  },
  "Muffin": {
    "count": 8,
    "lastSeenTime": 1586913798295
  },
  "Halbkreis": {
    "count": 3,
    "lastSeenTime": 1586899902236
  },
  "Schreibfeder": {
    "count": 2,
    "lastSeenTime": 1586816295247
  },
  "Tierarzt": {
    "count": 4,
    "lastSeenTime": 1586916441133
  },
  "Navy": {
    "count": 11,
    "lastSeenTime": 1586958897546
  },
  "Ferkel": {
    "count": 5,
    "lastSeenTime": 1586913467503,
    "difficulty": 1,
    "difficultyWeight": 5
  },
  "Patronenhülse": {
    "count": 6,
    "lastSeenTime": 1586873513174,
    "difficulty": 0.9166666666666666,
    "difficultyWeight": 12
  },
  "Bibel": {
    "count": 6,
    "lastSeenTime": 1586839953759
  },
  "geizig": {
    "count": 3,
    "lastSeenTime": 1586889892808
  },
  "Minze": {
    "count": 5,
    "lastSeenTime": 1586957855578
  },
  "Wolle": {
    "count": 8,
    "lastSeenTime": 1586916731964
  },
  "Cappuccino": {
    "count": 4,
    "lastSeenTime": 1586955517665
  },
  "Krankenwagen": {
    "count": 5,
    "lastSeenTime": 1586920982151
  },
  "Lüfter": {
    "count": 4,
    "lastSeenTime": 1586895847874
  },
  "Wal": {
    "count": 4,
    "lastSeenTime": 1586896080519,
    "difficulty": 1,
    "difficultyWeight": 3
  },
  "Ufer": {
    "count": 6,
    "lastSeenTime": 1586873419448
  },
  "chemisch": {
    "count": 12,
    "lastSeenTime": 1586959342619,
    "difficulty": 0.375,
    "difficultyWeight": 8
  },
  "Lisa Simpson": {
    "count": 10,
    "lastSeenTime": 1586903158426
  },
  "zurückspulen": {
    "count": 8,
    "lastSeenTime": 1586955300241
  },
  "Garfield": {
    "count": 8,
    "lastSeenTime": 1586869383090
  },
  "Schal": {
    "count": 7,
    "lastSeenTime": 1586958858535,
    "difficulty": 1,
    "difficultyWeight": 1
  },
  "Tischtennisschläger": {
    "count": 4,
    "lastSeenTime": 1586889489199
  },
  "Flagge": {
    "count": 5,
    "lastSeenTime": 1586915442369
  },
  "E-Gitarre": {
    "count": 7,
    "lastSeenTime": 1586903097380,
    "difficulty": 1,
    "difficultyWeight": 2
  },
  "einkaufen": {
    "count": 9,
    "lastSeenTime": 1586889566614,
    "difficulty": 0.8333333333333334,
    "difficultyWeight": 6
  },
  "Beerdigung": {
    "count": 5,
    "lastSeenTime": 1586904002117,
    "difficulty": 1,
    "difficultyWeight": 3
  },
  "Zylinder": {
    "count": 3,
    "lastSeenTime": 1586873715364
  },
  "Toast": {
    "count": 4,
    "lastSeenTime": 1586915478982
  },
  "Kompass": {
    "count": 5,
    "lastSeenTime": 1586832545576
  },
  "Schnellstraße": {
    "count": 8,
    "lastSeenTime": 1586919736561
  },
  "Sonnensystem": {
    "count": 6,
    "lastSeenTime": 1586955063993
  },
  "Liebe": {
    "count": 6,
    "lastSeenTime": 1586920421902
  },
  "Windsurfer": {
    "count": 5,
    "lastSeenTime": 1586920446574
  },
  "Fingerpuppe": {
    "count": 2,
    "lastSeenTime": 1586712445503
  },
  "Frisbee": {
    "count": 5,
    "lastSeenTime": 1586880408576,
    "difficulty": 1,
    "difficultyWeight": 1
  },
  "Rune": {
    "count": 6,
    "lastSeenTime": 1586915456714
  },
  "Mistgabel": {
    "count": 5,
    "lastSeenTime": 1586919664805
  },
  "Komiker": {
    "count": 4,
    "lastSeenTime": 1586915960185
  },
  "Floh": {
    "count": 6,
    "lastSeenTime": 1586870018417,
    "difficulty": 0.8571428571428571,
    "difficultyWeight": 7
  },
  "Pinienkerne": {
    "count": 1,
    "lastSeenTime": 1586580794838
  },
  "Puppe": {
    "count": 5,
    "lastSeenTime": 1586958154127
  },
  "Aktion": {
    "count": 5,
    "lastSeenTime": 1586896095108
  },
  "Amsel": {
    "count": 8,
    "lastSeenTime": 1586959738785
  },
  "Cockpit": {
    "count": 10,
    "lastSeenTime": 1586919477051,
    "difficulty": 0.7142857142857143,
    "difficultyWeight": 7
  },
  "Patenonkel": {
    "count": 7,
    "lastSeenTime": 1586915924921
  },
  "Sandkasten": {
    "count": 4,
    "lastSeenTime": 1586914806527
  },
  "Ampelmännchen": {
    "count": 4,
    "lastSeenTime": 1586955063993,
    "difficulty": 0.8823529411764706,
    "difficultyWeight": 34
  },
  "voll": {
    "count": 4,
    "lastSeenTime": 1586915806469
  },
  "Ostern": {
    "count": 3,
    "lastSeenTime": 1586879394962,
    "difficulty": 0.5,
    "difficultyWeight": 8
  },
  "feuerfest": {
    "count": 8,
    "lastSeenTime": 1586954774095
  },
  "Geige": {
    "count": 7,
    "lastSeenTime": 1586920872823,
    "difficulty": 0.8571428571428571,
    "difficultyWeight": 28
  },
  "Kiwi": {
    "count": 10,
    "lastSeenTime": 1586919501543,
    "publicGameCount": 1,
    "difficulty": 1,
    "difficultyWeight": 33
  },
  "Getränk": {
    "count": 7,
    "lastSeenTime": 1586919722216
  },
  "Ingenieur": {
    "count": 3,
    "lastSeenTime": 1586903281878,
    "difficulty": 1,
    "difficultyWeight": 3
  },
  "Tetris": {
    "count": 8,
    "lastSeenTime": 1586914932818
  },
  "Ast": {
    "count": 3,
    "lastSeenTime": 1586868143512,
    "difficulty": 1,
    "difficultyWeight": 38
  },
  "Dinosaurier": {
    "count": 4,
    "lastSeenTime": 1586915099808,
    "difficulty": 0.33333333333333337,
    "difficultyWeight": 9
  },
  "Spritze": {
    "count": 3,
    "lastSeenTime": 1586954763962,
    "publicGameCount": 1,
    "difficulty": 0.42857142857142855,
    "difficultyWeight": 7
  },
  "Schaukelstuhl": {
    "count": 7,
    "lastSeenTime": 1586881119264
  },
  "Google": {
    "count": 6,
    "lastSeenTime": 1586959443697
  },
  "Tukan": {
    "count": 6,
    "lastSeenTime": 1586900941854,
    "difficulty": 1,
    "difficultyWeight": 6
  },
  "Hackbraten": {
    "count": 4,
    "lastSeenTime": 1586919434514
  },
  "Trainer": {
    "count": 6,
    "lastSeenTime": 1586914465537
  },
  "arbeiten": {
    "count": 7,
    "lastSeenTime": 1586902658307
  },
  "Streichholz": {
    "count": 5,
    "lastSeenTime": 1586896595013,
    "difficulty": 1,
    "difficultyWeight": 20
  },
  "Lüftung": {
    "count": 2,
    "lastSeenTime": 1586817610330,
    "difficulty": 0.75,
    "difficultyWeight": 8
  },
  "Erfindung": {
    "count": 5,
    "lastSeenTime": 1586902587810
  },
  "Origami": {
    "count": 5,
    "lastSeenTime": 1586913757766
  },
  "Armband": {
    "count": 6,
    "lastSeenTime": 1586955324504
  },
  "Hügel": {
    "count": 5,
    "lastSeenTime": 1586920552243
  },
  "Bestatter": {
    "count": 10,
    "lastSeenTime": 1586874409884,
    "publicGameCount": 1,
    "difficulty": 0.8,
    "difficultyWeight": 10
  },
  "Wissenschaftler": {
    "count": 5,
    "lastSeenTime": 1586920324660
  },
  "Grafiktablett": {
    "count": 6,
    "lastSeenTime": 1586879298099,
    "difficulty": 0.8181818181818182,
    "difficultyWeight": 11
  },
  "Sand": {
    "count": 5,
    "lastSeenTime": 1586916153776
  },
  "Lady": {
    "count": 6,
    "lastSeenTime": 1586913893763
  },
  "Kazoo": {
    "count": 2,
    "lastSeenTime": 1586915737415
  },
  "toter Winkel": {
    "count": 6,
    "lastSeenTime": 1586891099195
  },
  "Omelett": {
    "count": 5,
    "lastSeenTime": 1586880936347
  },
  "Pluto": {
    "count": 3,
    "lastSeenTime": 1586859933316,
    "difficulty": 1,
    "difficultyWeight": 7
  },
  "Schornstein": {
    "count": 5,
    "lastSeenTime": 1586920265574,
    "difficulty": 1,
    "difficultyWeight": 7
  },
  "Mottenkugel": {
    "count": 5,
    "lastSeenTime": 1586955249129
  },
  "Zaubertrank": {
    "count": 5,
    "lastSeenTime": 1586958027091
  },
  "Gully": {
    "count": 2,
    "lastSeenTime": 1586955838030
  },
  "Pobacken": {
    "count": 3,
    "lastSeenTime": 1586908597667,
    "difficulty": 1,
    "difficultyWeight": 1
  },
  "Stewardess": {
    "count": 5,
    "lastSeenTime": 1586907472056
  },
  "Klaviersaite": {
    "count": 8,
    "lastSeenTime": 1586920895382
  },
  "Organ": {
    "count": 9,
    "lastSeenTime": 1586959767498
  },
  "Fußball": {
    "count": 7,
    "lastSeenTime": 1586901146209
  },
  "Bedienung": {
    "count": 3,
    "lastSeenTime": 1586909030886
  },
  "Grafik": {
    "count": 3,
    "lastSeenTime": 1586920153587
  },
  "Teufel": {
    "count": 6,
    "lastSeenTime": 1586818535737,
    "publicGameCount": 1,
    "difficulty": 1,
    "difficultyWeight": 11
  },
  "Linie": {
    "count": 6,
    "lastSeenTime": 1586890443871
  },
  "Jupiter": {
    "count": 5,
    "lastSeenTime": 1586860123492
  },
  "Lanze": {
    "count": 10,
    "lastSeenTime": 1586913962978
  },
  "Ring": {
    "count": 2,
    "lastSeenTime": 1586724556569,
    "publicGameCount": 1,
    "difficulty": 0.6538461538461539,
    "difficultyWeight": 26
  },
  "High Five": {
    "count": 6,
    "lastSeenTime": 1586896166029,
    "publicGameCount": 1,
    "difficulty": 0.8235294117647058,
    "difficultyWeight": 17
  },
  "Wimper": {
    "count": 6,
    "lastSeenTime": 1586913497512,
    "difficulty": 0.9428571428571428,
    "difficultyWeight": 35
  },
  "zoomen": {
    "count": 5,
    "lastSeenTime": 1586921094172
  },
  "Rechen": {
    "count": 4,
    "lastSeenTime": 1586958253918
  },
  "Amboss": {
    "count": 4,
    "lastSeenTime": 1586889697457,
    "difficulty": 1,
    "difficultyWeight": 7
  },
  "Zickzack": {
    "count": 12,
    "lastSeenTime": 1586954688633
  },
  "Ohrhörer": {
    "count": 7,
    "lastSeenTime": 1586914243836,
    "difficulty": 1,
    "difficultyWeight": 3
  },
  "Parfüm": {
    "count": 3,
    "lastSeenTime": 1586915021163
  },
  "Rollschuhe": {
    "count": 6,
    "lastSeenTime": 1586908103784,
    "difficulty": 0.8636363636363636,
    "difficultyWeight": 44
  },
  "Hippie": {
    "count": 4,
    "lastSeenTime": 1586718322766
  },
  "Kuba": {
    "count": 5,
    "lastSeenTime": 1586900776963
  },
  "Rolltreppe": {
    "count": 9,
    "lastSeenTime": 1586903048902
  },
  "Schlange": {
    "count": 7,
    "lastSeenTime": 1586920751201
  },
  "Ellbogen": {
    "count": 4,
    "lastSeenTime": 1586904161945,
    "difficulty": 1,
    "difficultyWeight": 1
  },
  "Donut": {
    "count": 6,
    "lastSeenTime": 1586958289143,
    "difficulty": 1,
    "difficultyWeight": 21
  },
  "Zeppelin": {
    "count": 10,
    "lastSeenTime": 1586955419506,
    "difficulty": 1,
    "difficultyWeight": 46
  },
  "reinigen": {
    "count": 6,
    "lastSeenTime": 1586858960624
  },
  "Wasserpfeife": {
    "count": 4,
    "lastSeenTime": 1586916637868
  },
  "beizen": {
    "count": 8,
    "lastSeenTime": 1586908996510
  },
  "Comicbuch": {
    "count": 5,
    "lastSeenTime": 1586915624112,
    "difficulty": 1,
    "difficultyWeight": 10
  },
  "Palast": {
    "count": 7,
    "lastSeenTime": 1586908204031
  },
  "Yoda": {
    "count": 3,
    "lastSeenTime": 1586955010181,
    "difficulty": 1,
    "difficultyWeight": 2
  },
  "Mixer": {
    "count": 6,
    "lastSeenTime": 1586915339348,
    "publicGameCount": 1,
    "difficulty": 0.5,
    "difficultyWeight": 8
  },
  "Hosenstall": {
    "count": 4,
    "lastSeenTime": 1586956336636,
    "publicGameCount": 1,
    "difficulty": 0.5714285714285714,
    "difficultyWeight": 7
  },
  "Chinchilla": {
    "count": 4,
    "lastSeenTime": 1586908945942
  },
  "Briefkasten": {
    "count": 5,
    "lastSeenTime": 1586881205453
  },
  "The Beatles": {
    "count": 6,
    "lastSeenTime": 1586959232575
  },
  "Birke": {
    "count": 8,
    "lastSeenTime": 1586902687551
  },
  "Brunnen": {
    "count": 4,
    "lastSeenTime": 1586826911367
  },
  "Reifen": {
    "count": 5,
    "lastSeenTime": 1586873308659
  },
  "Dosenöffner": {
    "count": 6,
    "lastSeenTime": 1586891228286
  },
  "Ordner": {
    "count": 8,
    "lastSeenTime": 1586958278719,
    "difficulty": 1,
    "difficultyWeight": 3
  },
  "Müllabfuhr": {
    "count": 3,
    "lastSeenTime": 1586900240367
  },
  "Mais": {
    "count": 3,
    "lastSeenTime": 1586900866878,
    "publicGameCount": 1,
    "difficulty": 0.8,
    "difficultyWeight": 10
  },
  "Galaxie": {
    "count": 6,
    "lastSeenTime": 1586907754016,
    "publicGameCount": 1,
    "difficulty": 0.75,
    "difficultyWeight": 4
  },
  "Tweety": {
    "count": 9,
    "lastSeenTime": 1586908757151
  },
  "Kirschblüte": {
    "count": 5,
    "lastSeenTime": 1586920690415
  },
  "Anhänger": {
    "count": 4,
    "lastSeenTime": 1586916074772
  },
  "Milchshake": {
    "count": 6,
    "lastSeenTime": 1586913988149
  },
  "Irokesenfrisur": {
    "count": 5,
    "lastSeenTime": 1586958944806,
    "difficulty": 0.6666666666666667,
    "difficultyWeight": 9
  },
  "Sohn": {
    "count": 1,
    "lastSeenTime": 1586589989733,
    "difficulty": 1,
    "difficultyWeight": 1
  },
  "Schneeballschlacht": {
    "count": 3,
    "lastSeenTime": 1586954910317
  },
  "sabbern": {
    "count": 6,
    "lastSeenTime": 1586954939261
  },
  "Tintenpatrone": {
    "count": 5,
    "lastSeenTime": 1586959142159
  },
  "Essig": {
    "count": 7,
    "lastSeenTime": 1586907754016
  },
  "Lilie": {
    "count": 6,
    "lastSeenTime": 1586955370991,
    "difficulty": 1,
    "difficultyWeight": 6
  },
  "Ferse": {
    "count": 3,
    "lastSeenTime": 1586916540283,
    "difficulty": 0.42857142857142855,
    "difficultyWeight": 7
  },
  "föhnen": {
    "count": 7,
    "lastSeenTime": 1586959131863
  },
  "Akne": {
    "count": 4,
    "lastSeenTime": 1586880759383
  },
  "Prinz": {
    "count": 6,
    "lastSeenTime": 1586889638222
  },
  "Ampel": {
    "count": 6,
    "lastSeenTime": 1586889591248,
    "publicGameCount": 2,
    "difficulty": 0.7333333333333333,
    "difficultyWeight": 45
  },
  "Patrick Star": {
    "count": 8,
    "lastSeenTime": 1586954967541,
    "publicGameCount": 1,
    "difficulty": 0.875,
    "difficultyWeight": 8
  },
  "Einrad": {
    "count": 6,
    "lastSeenTime": 1586890806055,
    "difficulty": 0.9,
    "difficultyWeight": 10
  },
  "Schwalbe": {
    "count": 5,
    "lastSeenTime": 1586901274634
  },
  "Postbote": {
    "count": 6,
    "lastSeenTime": 1586954603173,
    "difficulty": 0.7,
    "difficultyWeight": 10
  },
  "Korkenzieher": {
    "count": 3,
    "lastSeenTime": 1586814402764
  },
  "Heiligenschein": {
    "count": 6,
    "lastSeenTime": 1586859642077
  },
  "Rennauto": {
    "count": 3,
    "lastSeenTime": 1586874409884,
    "difficulty": 0.9,
    "difficultyWeight": 10
  },
  "Schüssel": {
    "count": 6,
    "lastSeenTime": 1586959207698,
    "difficulty": 1,
    "difficultyWeight": 1
  },
  "Leck": {
    "count": 6,
    "lastSeenTime": 1586958668588
  },
  "prallen": {
    "count": 6,
    "lastSeenTime": 1586916493492,
    "difficulty": 1,
    "difficultyWeight": 36
  },
  "Schweißer": {
    "count": 8,
    "lastSeenTime": 1586921283836
  },
  "Köcher": {
    "count": 7,
    "lastSeenTime": 1586958239433
  },
  "Rote Beete": {
    "count": 6,
    "lastSeenTime": 1586920237009
  },
  "Hollywood": {
    "count": 4,
    "lastSeenTime": 1586880759383
  },
  "Bob Ross": {
    "count": 6,
    "lastSeenTime": 1586914515993
  },
  "Donnerstag": {
    "count": 3,
    "lastSeenTime": 1586826344267,
    "difficulty": 0.4117647058823529,
    "difficultyWeight": 17
  },
  "Fallschirm": {
    "count": 9,
    "lastSeenTime": 1586908583500,
    "difficulty": 1,
    "difficultyWeight": 11
  },
  "Gorilla": {
    "count": 7,
    "lastSeenTime": 1586956033779
  },
  "Orchester": {
    "count": 6,
    "lastSeenTime": 1586956336636
  },
  "Lagerhaus": {
    "count": 5,
    "lastSeenTime": 1586859907934
  },
  "Bienenkönigin": {
    "count": 4,
    "lastSeenTime": 1586901397893,
    "difficulty": 0.7878787878787878,
    "difficultyWeight": 33
  },
  "Mayonnaise": {
    "count": 5,
    "lastSeenTime": 1586955716758
  },
  "Reptil": {
    "count": 12,
    "lastSeenTime": 1586915059551
  },
  "Lichtschwert": {
    "count": 4,
    "lastSeenTime": 1586896127512,
    "difficulty": 1,
    "difficultyWeight": 6
  },
  "Sandburg": {
    "count": 3,
    "lastSeenTime": 1586807745991
  },
  "Brathering": {
    "count": 2,
    "lastSeenTime": 1586895949288
  },
  "Leonardo da Vinci": {
    "count": 5,
    "lastSeenTime": 1586900226190
  },
  "U-Bahn": {
    "count": 6,
    "lastSeenTime": 1586916574944
  },
  "Echo": {
    "count": 5,
    "lastSeenTime": 1586897250506,
    "publicGameCount": 1,
    "difficulty": 0.8333333333333334,
    "difficultyWeight": 6
  },
  "Biss": {
    "count": 1,
    "lastSeenTime": 1586590631736
  },
  "Fernbedienung": {
    "count": 3,
    "lastSeenTime": 1586913648704
  },
  "Vodka": {
    "count": 6,
    "lastSeenTime": 1586919711626
  },
  "Captain America": {
    "count": 4,
    "lastSeenTime": 1586916038112
  },
  "Pfeffer": {
    "count": 7,
    "lastSeenTime": 1586920926700,
    "difficulty": 0.7428571428571429,
    "difficultyWeight": 35,
    "publicGameCount": 1
  },
  "Gummibärchen": {
    "count": 4,
    "lastSeenTime": 1586916063451
  },
  "Beethoven": {
    "count": 4,
    "lastSeenTime": 1586954900699
  },
  "Rasierklinge": {
    "count": 7,
    "lastSeenTime": 1586958264090
  },
  "Seiltänzer": {
    "count": 5,
    "lastSeenTime": 1586900749260
  },
  "Zucchini": {
    "count": 5,
    "lastSeenTime": 1586955542167,
    "difficulty": 0.6666666666666666,
    "difficultyWeight": 24
  },
  "Sandsturm": {
    "count": 3,
    "lastSeenTime": 1586959547988
  },
  "Fliegenpilz": {
    "count": 8,
    "lastSeenTime": 1586881220427,
    "difficulty": 0.8888888888888888,
    "difficultyWeight": 9
  },
  "Zügel": {
    "count": 7,
    "lastSeenTime": 1586915176554
  },
  "Barack Obama": {
    "count": 7,
    "lastSeenTime": 1586839026220
  },
  "Schleife": {
    "count": 4,
    "lastSeenTime": 1586913868342
  },
  "Mahlzeit": {
    "count": 6,
    "lastSeenTime": 1586880019286,
    "difficulty": 1,
    "difficultyWeight": 1
  },
  "Autoradio": {
    "count": 8,
    "lastSeenTime": 1586902787660,
    "publicGameCount": 1,
    "difficulty": 0.75,
    "difficultyWeight": 4
  },
  "lernen": {
    "count": 5,
    "lastSeenTime": 1586897039497
  },
  "Matsch": {
    "count": 3,
    "lastSeenTime": 1586815745247
  },
  "seekrank": {
    "count": 3,
    "lastSeenTime": 1586908382794,
    "difficulty": 0.25,
    "difficultyWeight": 8
  },
  "Olaf": {
    "count": 6,
    "lastSeenTime": 1586915919352
  },
  "Schaf": {
    "count": 8,
    "lastSeenTime": 1586890490289,
    "publicGameCount": 1
  },
  "Spiegel": {
    "count": 4,
    "lastSeenTime": 1586881220427,
    "difficulty": 1,
    "difficultyWeight": 14
  },
  "positiv": {
    "count": 3,
    "lastSeenTime": 1586956363582,
    "difficulty": 0.8305084745762712,
    "difficultyWeight": 59
  },
  "Cheeseburger": {
    "count": 5,
    "lastSeenTime": 1586958561895
  },
  "Liechtenstein": {
    "count": 5,
    "lastSeenTime": 1586896569941
  },
  "fabelhaft": {
    "count": 5,
    "lastSeenTime": 1586909006610
  },
  "Winzer": {
    "count": 2,
    "lastSeenTime": 1586908395891
  },
  "kopieren": {
    "count": 7,
    "lastSeenTime": 1586959242580,
    "difficulty": 1,
    "difficultyWeight": 50
  },
  "Doritos": {
    "count": 6,
    "lastSeenTime": 1586919423610
  },
  "gehorchen": {
    "count": 6,
    "lastSeenTime": 1586920836680
  },
  "Niederlande": {
    "count": 2,
    "lastSeenTime": 1586903158426
  },
  "WhatsApp": {
    "count": 7,
    "lastSeenTime": 1586914491372,
    "difficulty": 0.5,
    "difficultyWeight": 4
  },
  "Fernseher": {
    "count": 4,
    "lastSeenTime": 1586919811805,
    "difficulty": 0.9333333333333333,
    "difficultyWeight": 15
  },
  "Geld": {
    "count": 2,
    "lastSeenTime": 1586726008015
  },
  "Wrestling": {
    "count": 5,
    "lastSeenTime": 1586900650119
  },
  "tot": {
    "count": 11,
    "lastSeenTime": 1586958212167
  },
  "Kuhglocke": {
    "count": 10,
    "lastSeenTime": 1586901321700
  },
  "Gehirn": {
    "count": 7,
    "lastSeenTime": 1586958859906,
    "difficulty": 1,
    "difficultyWeight": 5
  },
  "New York": {
    "count": 9,
    "lastSeenTime": 1586914373601
  },
  "Gepäck": {
    "count": 10,
    "lastSeenTime": 1586956016971,
    "difficulty": 1,
    "difficultyWeight": 16
  },
  "Spartacus": {
    "count": 7,
    "lastSeenTime": 1586890454038
  },
  "Kabel": {
    "count": 7,
    "lastSeenTime": 1586806153755
  },
  "Armaturenbrett": {
    "count": 5,
    "lastSeenTime": 1586954674416
  },
  "Wein": {
    "count": 7,
    "lastSeenTime": 1586955923939,
    "difficulty": 0.36363636363636365,
    "difficultyWeight": 11
  },
  "Michael Jackson": {
    "count": 3,
    "lastSeenTime": 1586920375356
  },
  "Videospiel": {
    "count": 9,
    "lastSeenTime": 1586913482555
  },
  "Curry": {
    "count": 10,
    "lastSeenTime": 1586955532027
  },
  "Kitesurfen": {
    "count": 1,
    "lastSeenTime": 1586591460245
  },
  "Reis": {
    "count": 8,
    "lastSeenTime": 1586907764126
  },
  "Jenga": {
    "count": 9,
    "lastSeenTime": 1586915089601
  },
  "massieren": {
    "count": 3,
    "lastSeenTime": 1586840060848
  },
  "Zündschnur": {
    "count": 4,
    "lastSeenTime": 1586954999969,
    "difficulty": 1,
    "difficultyWeight": 7
  },
  "Eigelb": {
    "count": 4,
    "lastSeenTime": 1586708871882
  },
  "Wurm": {
    "count": 7,
    "lastSeenTime": 1586858975693
  },
  "unsichtbar": {
    "count": 7,
    "lastSeenTime": 1586920922732
  },
  "gut": {
    "count": 4,
    "lastSeenTime": 1586817823665
  },
  "Schottenrock": {
    "count": 5,
    "lastSeenTime": 1586900254522
  },
  "Kopfende": {
    "count": 5,
    "lastSeenTime": 1586903572367,
    "difficulty": 1,
    "difficultyWeight": 6
  },
  "Fischgräte": {
    "count": 4,
    "lastSeenTime": 1586833356919
  },
  "nichts": {
    "count": 5,
    "lastSeenTime": 1586955823724
  },
  "Braut": {
    "count": 14,
    "lastSeenTime": 1586916680436
  },
  "Mischpult": {
    "count": 6,
    "lastSeenTime": 1586860406384
  },
  "Lehm": {
    "count": 7,
    "lastSeenTime": 1586889638222
  },
  "Erdnuss": {
    "count": 7,
    "lastSeenTime": 1586958887354
  },
  "Briefträger": {
    "count": 4,
    "lastSeenTime": 1586959269543,
    "difficulty": 0.9375,
    "difficultyWeight": 16
  },
  "Miniclip": {
    "count": 7,
    "lastSeenTime": 1586913998392
  },
  "Blumenstrauß": {
    "count": 4,
    "lastSeenTime": 1586916013387
  },
  "Ziegelstein": {
    "count": 6,
    "lastSeenTime": 1586840420314
  },
  "Öl": {
    "count": 4,
    "lastSeenTime": 1586958212167
  },
  "Aprikose": {
    "count": 5,
    "lastSeenTime": 1586891163944,
    "publicGameCount": 1,
    "difficulty": 0.6666666666666667,
    "difficultyWeight": 9
  },
  "Taube": {
    "count": 6,
    "lastSeenTime": 1586901052765
  },
  "Wissenschaft": {
    "count": 6,
    "lastSeenTime": 1586916237094
  },
  "Champagnersorbet": {
    "count": 6,
    "lastSeenTime": 1586896666647
  },
  "fallen": {
    "count": 2,
    "lastSeenTime": 1586805561130
  },
  "Israel": {
    "count": 4,
    "lastSeenTime": 1586827800087,
    "difficulty": 1,
    "difficultyWeight": 4
  },
  "Milchstraße": {
    "count": 8,
    "lastSeenTime": 1586900660239
  },
  "Wippe": {
    "count": 4,
    "lastSeenTime": 1586955726995,
    "difficulty": 1,
    "difficultyWeight": 5
  },
  "Kaviar": {
    "count": 6,
    "lastSeenTime": 1586959622590
  },
  "Hafen": {
    "count": 7,
    "lastSeenTime": 1586958171361,
    "publicGameCount": 1,
    "difficulty": 0.7931034482758621,
    "difficultyWeight": 29
  },
  "Minute": {
    "count": 8,
    "lastSeenTime": 1586897516933
  },
  "Zuckerguss": {
    "count": 2,
    "lastSeenTime": 1586707490326
  },
  "Bock": {
    "count": 3,
    "lastSeenTime": 1586806457936
  },
  "Fleisch": {
    "count": 6,
    "lastSeenTime": 1586903122787,
    "difficulty": 0.875,
    "difficultyWeight": 8
  },
  "Schnabel": {
    "count": 7,
    "lastSeenTime": 1586958456673
  },
  "Erdbeben": {
    "count": 5,
    "lastSeenTime": 1586900724985
  },
  "Zwerg": {
    "count": 1,
    "lastSeenTime": 1586595084509
  },
  "Außerirdischer": {
    "count": 4,
    "lastSeenTime": 1586902687551
  },
  "Zunge": {
    "count": 4,
    "lastSeenTime": 1586920922732
  },
  "Harz": {
    "count": 5,
    "lastSeenTime": 1586915312713
  },
  "Nasa": {
    "count": 5,
    "lastSeenTime": 1586896902340
  },
  "LKW": {
    "count": 7,
    "lastSeenTime": 1586956009184
  },
  "Salz": {
    "count": 6,
    "lastSeenTime": 1586896095108
  },
  "Schicksal": {
    "count": 7,
    "lastSeenTime": 1586920326688
  },
  "arm": {
    "count": 7,
    "lastSeenTime": 1586919522211,
    "difficulty": 1,
    "difficultyWeight": 1
  },
  "Schmetterling": {
    "count": 4,
    "lastSeenTime": 1586958289143
  },
  "Mäusefalle": {
    "count": 10,
    "lastSeenTime": 1586956504597,
    "difficulty": 0.9444444444444444,
    "difficultyWeight": 36
  },
  "Hals": {
    "count": 3,
    "lastSeenTime": 1586826301547
  },
  "Teddybär": {
    "count": 6,
    "lastSeenTime": 1586955370991,
    "difficulty": 0.8235294117647058,
    "difficultyWeight": 17
  },
  "Biber": {
    "count": 7,
    "lastSeenTime": 1586916461517,
    "difficulty": 1,
    "difficultyWeight": 16
  },
  "Eminem": {
    "count": 3,
    "lastSeenTime": 1586904308260
  },
  "Nasenring": {
    "count": 6,
    "lastSeenTime": 1586955823724,
    "difficulty": 1,
    "difficultyWeight": 9
  },
  "Daffy Duck": {
    "count": 4,
    "lastSeenTime": 1586873936206
  },
  "Darm": {
    "count": 8,
    "lastSeenTime": 1586959255115
  },
  "kleiner Finger": {
    "count": 2,
    "lastSeenTime": 1586651151122,
    "publicGameCount": 1,
    "difficulty": 0.4666666666666667,
    "difficultyWeight": 15
  },
  "Mütze": {
    "count": 3,
    "lastSeenTime": 1586815197615,
    "difficulty": 1,
    "difficultyWeight": 8
  },
  "Sitzbank": {
    "count": 9,
    "lastSeenTime": 1586958471219,
    "publicGameCount": 1,
    "difficulty": 0.9655172413793104,
    "difficultyWeight": 29
  },
  "Schiffswrack": {
    "count": 2,
    "lastSeenTime": 1586596602460
  },
  "Archäologe": {
    "count": 3,
    "lastSeenTime": 1586840929412
  },
  "Duell": {
    "count": 8,
    "lastSeenTime": 1586907557164,
    "difficulty": 0.5,
    "difficultyWeight": 12
  },
  "Kröte": {
    "count": 8,
    "lastSeenTime": 1586908626057
  },
  "Punkte": {
    "count": 3,
    "lastSeenTime": 1586904287692,
    "difficulty": 0.6666666666666666,
    "difficultyWeight": 3
  },
  "Kojote": {
    "count": 6,
    "lastSeenTime": 1586956570180
  },
  "Reddit": {
    "count": 9,
    "lastSeenTime": 1586958052113,
    "difficulty": 0.8333333333333334,
    "difficultyWeight": 18
  },
  "iPad": {
    "count": 6,
    "lastSeenTime": 1586896714278
  },
  "Katapult": {
    "count": 3,
    "lastSeenTime": 1586919983033
  },
  "Leiterwagen": {
    "count": 6,
    "lastSeenTime": 1586901191898
  },
  "Schalter": {
    "count": 6,
    "lastSeenTime": 1586833150052
  },
  "Lotterie": {
    "count": 3,
    "lastSeenTime": 1586955644811,
    "difficulty": 0.9583333333333334,
    "difficultyWeight": 24
  },
  "China": {
    "count": 3,
    "lastSeenTime": 1586915856980,
    "difficulty": 1,
    "difficultyWeight": 7
  },
  "Anführer": {
    "count": 4,
    "lastSeenTime": 1586815327708
  },
  "Waschstraße": {
    "count": 4,
    "lastSeenTime": 1586874562516
  },
  "Seetang": {
    "count": 4,
    "lastSeenTime": 1586955542167,
    "difficulty": 1,
    "difficultyWeight": 11
  },
  "Tasmanischer Teufel": {
    "count": 3,
    "lastSeenTime": 1586916368116
  },
  "bauchpinseln": {
    "count": 5,
    "lastSeenTime": 1586879160227,
    "publicGameCount": 1
  },
  "Filter": {
    "count": 3,
    "lastSeenTime": 1586859018427
  },
  "Bett": {
    "count": 9,
    "lastSeenTime": 1586955409372,
    "difficulty": 1,
    "difficultyWeight": 7
  },
  "Schlüsselanhänger": {
    "count": 4,
    "lastSeenTime": 1586916461517,
    "difficulty": 1,
    "difficultyWeight": 6
  },
  "Giraffe": {
    "count": 2,
    "lastSeenTime": 1586874945892,
    "difficulty": 1,
    "difficultyWeight": 9
  },
  "Wildschwein": {
    "count": 6,
    "lastSeenTime": 1586956130879
  },
  "Kino": {
    "count": 3,
    "lastSeenTime": 1586921170819
  },
  "Chirurg": {
    "count": 9,
    "lastSeenTime": 1586889475038,
    "difficulty": 0.8,
    "difficultyWeight": 5
  },
  "Lampenschirm": {
    "count": 6,
    "lastSeenTime": 1586914094806
  },
  "schwach": {
    "count": 6,
    "lastSeenTime": 1586889917123
  },
  "Leonardo DiCaprio": {
    "count": 4,
    "lastSeenTime": 1586827438446
  },
  "Pringles": {
    "count": 2,
    "lastSeenTime": 1586915417224
  },
  "Terrasse": {
    "count": 5,
    "lastSeenTime": 1586840254824,
    "difficulty": 0.8787878787878788,
    "difficultyWeight": 33
  },
  "rauchen": {
    "count": 7,
    "lastSeenTime": 1586873988685
  },
  "Kerze": {
    "count": 2,
    "lastSeenTime": 1586717755995,
    "difficulty": 0.7,
    "difficultyWeight": 10
  },
  "gelb": {
    "count": 3,
    "lastSeenTime": 1586908541127
  },
  "Bitcoin": {
    "count": 5,
    "lastSeenTime": 1586920180037
  },
  "Homer Simpson": {
    "count": 6,
    "lastSeenTime": 1586840216366
  },
  "Patrone": {
    "count": 4,
    "lastSeenTime": 1586889725951
  },
  "Kuh": {
    "count": 7,
    "lastSeenTime": 1586956324955
  },
  "Mülleimer": {
    "count": 11,
    "lastSeenTime": 1586959020405
  },
  "Sushi": {
    "count": 6,
    "lastSeenTime": 1586896805974
  },
  "Poker": {
    "count": 3,
    "lastSeenTime": 1586859615003
  },
  "Waffel": {
    "count": 7,
    "lastSeenTime": 1586958794944
  },
  "Gazelle": {
    "count": 4,
    "lastSeenTime": 1586908518895,
    "difficulty": 0.7857142857142857,
    "difficultyWeight": 14
  },
  "Morty": {
    "count": 2,
    "lastSeenTime": 1586785081909
  },
  "Waldbrand": {
    "count": 7,
    "lastSeenTime": 1586959344881
  },
  "Delle": {
    "count": 3,
    "lastSeenTime": 1586805672921
  },
  "Virus": {
    "count": 5,
    "lastSeenTime": 1586907788446
  },
  "Veganer": {
    "count": 3,
    "lastSeenTime": 1586868461473,
    "difficulty": 0.7142857142857143,
    "difficultyWeight": 21
  },
  "nähen": {
    "count": 4,
    "lastSeenTime": 1586727015282
  },
  "Sandwich": {
    "count": 5,
    "lastSeenTime": 1586901361357
  },
  "Ozean": {
    "count": 6,
    "lastSeenTime": 1586956570180,
    "difficulty": 0.5555555555555556,
    "difficultyWeight": 9
  },
  "Stern": {
    "count": 5,
    "lastSeenTime": 1586901461723
  },
  "Jahreszeit": {
    "count": 8,
    "lastSeenTime": 1586958499761
  },
  "Mozart": {
    "count": 5,
    "lastSeenTime": 1586832383649,
    "difficulty": 0.8571428571428571,
    "difficultyWeight": 7
  },
  "Schornsteinfeger": {
    "count": 7,
    "lastSeenTime": 1586879670490,
    "difficulty": 1,
    "difficultyWeight": 9
  },
  "Magazin": {
    "count": 5,
    "lastSeenTime": 1586916415277,
    "difficulty": 0.6666666666666666,
    "difficultyWeight": 6
  },
  "schmelzen": {
    "count": 6,
    "lastSeenTime": 1586890729326
  },
  "Silvester": {
    "count": 7,
    "lastSeenTime": 1586958375030,
    "difficulty": 0.9411764705882353,
    "difficultyWeight": 17
  },
  "Eintopf": {
    "count": 6,
    "lastSeenTime": 1586921071907
  },
  "Lidschatten": {
    "count": 5,
    "lastSeenTime": 1586955503431
  },
  "Traum": {
    "count": 5,
    "lastSeenTime": 1586959804291
  },
  "Kruste": {
    "count": 4,
    "lastSeenTime": 1586920432273
  },
  "pfeifen": {
    "count": 3,
    "lastSeenTime": 1586896458746,
    "difficulty": 0.4,
    "difficultyWeight": 10
  },
  "zertrümmern": {
    "count": 7,
    "lastSeenTime": 1586889393401
  },
  "chaotisch": {
    "count": 8,
    "lastSeenTime": 1586908175585
  },
  "Schmerz": {
    "count": 5,
    "lastSeenTime": 1586920119062
  },
  "Limbo": {
    "count": 4,
    "lastSeenTime": 1586879298099,
    "difficulty": 0.7684210526315789,
    "difficultyWeight": 95
  },
  "Baumwolle": {
    "count": 3,
    "lastSeenTime": 1586915099808
  },
  "Schneebesen": {
    "count": 5,
    "lastSeenTime": 1586813662034
  },
  "Balkon": {
    "count": 4,
    "lastSeenTime": 1586896769426,
    "difficulty": 0.5,
    "difficultyWeight": 8
  },
  "Freibad": {
    "count": 7,
    "lastSeenTime": 1586913827190
  },
  "kriechen": {
    "count": 7,
    "lastSeenTime": 1586958887354
  },
  "Model": {
    "count": 5,
    "lastSeenTime": 1586889489199
  },
  "Grillen": {
    "count": 7,
    "lastSeenTime": 1586956048018
  },
  "Schreibtisch": {
    "count": 9,
    "lastSeenTime": 1586959290196
  },
  "Leder": {
    "count": 4,
    "lastSeenTime": 1586920361074
  },
  "Osten": {
    "count": 2,
    "lastSeenTime": 1586634009548
  },
  "Rezeption": {
    "count": 6,
    "lastSeenTime": 1586913827190
  },
  "Madagaskar": {
    "count": 2,
    "lastSeenTime": 1586914309641
  },
  "rutschen": {
    "count": 7,
    "lastSeenTime": 1586919811805,
    "difficulty": 1,
    "difficultyWeight": 12
  },
  "Punkt": {
    "count": 3,
    "lastSeenTime": 1586868381050
  },
  "Frosch": {
    "count": 4,
    "lastSeenTime": 1586814181157
  },
  "Tennisschläger": {
    "count": 3,
    "lastSeenTime": 1586954603173
  },
  "Tragfläche": {
    "count": 3,
    "lastSeenTime": 1586889719019
  },
  "spielen": {
    "count": 4,
    "lastSeenTime": 1586904147395
  },
  "Australien": {
    "count": 3,
    "lastSeenTime": 1586817485337
  },
  "Schwan": {
    "count": 4,
    "lastSeenTime": 1586868119871
  },
  "Seekuh": {
    "count": 2,
    "lastSeenTime": 1586817081429
  },
  "stampfen": {
    "count": 5,
    "lastSeenTime": 1586900361272
  },
  "Vegetarier": {
    "count": 2,
    "lastSeenTime": 1586713274267
  },
  "Streichholzschachtel": {
    "count": 4,
    "lastSeenTime": 1586913648704,
    "publicGameCount": 1,
    "difficulty": 1,
    "difficultyWeight": 4
  },
  "Eimer": {
    "count": 5,
    "lastSeenTime": 1586959633128,
    "difficulty": 0.6,
    "difficultyWeight": 10
  },
  "Telefon": {
    "count": 5,
    "lastSeenTime": 1586919434514,
    "difficulty": 0.65,
    "difficultyWeight": 20
  },
  "Onkel": {
    "count": 4,
    "lastSeenTime": 1586897484538
  },
  "Warnweste": {
    "count": 9,
    "lastSeenTime": 1586916013387,
    "difficulty": 1,
    "difficultyWeight": 6
  },
  "Dompteur": {
    "count": 9,
    "lastSeenTime": 1586908931787
  },
  "Mittwoch": {
    "count": 9,
    "lastSeenTime": 1586955113604
  },
  "pink": {
    "count": 6,
    "lastSeenTime": 1586914080489,
    "difficulty": 0.5,
    "difficultyWeight": 12
  },
  "Krake": {
    "count": 4,
    "lastSeenTime": 1586958944806
  },
  "goldener Apfel": {
    "count": 7,
    "lastSeenTime": 1586895761615
  },
  "Maisfeld": {
    "count": 7,
    "lastSeenTime": 1586903711780
  },
  "Holzfäller": {
    "count": 2,
    "lastSeenTime": 1586727362775,
    "difficulty": 0.6571428571428571,
    "difficultyWeight": 35
  },
  "Familie": {
    "count": 6,
    "lastSeenTime": 1586920538046
  },
  "Paintball": {
    "count": 1,
    "lastSeenTime": 1586608344685
  },
  "Kuchen": {
    "count": 4,
    "lastSeenTime": 1586955263523
  },
  "Tauchgang": {
    "count": 7,
    "lastSeenTime": 1586904234099
  },
  "Filmriss": {
    "count": 5,
    "lastSeenTime": 1586915731157
  },
  "Kristall": {
    "count": 4,
    "lastSeenTime": 1586896530203
  },
  "Straßensperre": {
    "count": 3,
    "lastSeenTime": 1586859787263
  },
  "Jagdziel": {
    "count": 6,
    "lastSeenTime": 1586914258463,
    "publicGameCount": 1
  },
  "Zimmermann": {
    "count": 5,
    "lastSeenTime": 1586914792295
  },
  "Erkältung": {
    "count": 8,
    "lastSeenTime": 1586959344881
  },
  "Blase": {
    "count": 5,
    "lastSeenTime": 1586919413397
  },
  "Moos": {
    "count": 4,
    "lastSeenTime": 1586900318835,
    "difficulty": 1,
    "difficultyWeight": 7
  },
  "Einsiedler": {
    "count": 4,
    "lastSeenTime": 1586896040899,
    "publicGameCount": 1
  },
  "Kanada": {
    "count": 2,
    "lastSeenTime": 1586710340067,
    "difficulty": 0.8333333333333334,
    "difficultyWeight": 6
  },
  "Glühwürmchen": {
    "count": 3,
    "lastSeenTime": 1586833885395
  },
  "Sonnenschirm": {
    "count": 7,
    "lastSeenTime": 1586914428482
  },
  "Brieföffner": {
    "count": 7,
    "lastSeenTime": 1586957925802
  },
  "Orgel": {
    "count": 5,
    "lastSeenTime": 1586859816236
  },
  "Seifenblase": {
    "count": 4,
    "lastSeenTime": 1586920043786,
    "difficulty": 0.5714285714285714,
    "difficultyWeight": 7
  },
  "Notenschlüssel": {
    "count": 5,
    "lastSeenTime": 1586916118321
  },
  "Kroatien": {
    "count": 10,
    "lastSeenTime": 1586959728501
  },
  "Snowboard": {
    "count": 4,
    "lastSeenTime": 1586954577957
  },
  "Zahn": {
    "count": 3,
    "lastSeenTime": 1586890131974,
    "difficulty": 1,
    "difficultyWeight": 31
  },
  "Zugbrücke": {
    "count": 5,
    "lastSeenTime": 1586916717644
  },
  "Distanz": {
    "count": 4,
    "lastSeenTime": 1586907361110
  },
  "Dreieck": {
    "count": 9,
    "lastSeenTime": 1586955681359,
    "difficulty": 0.6,
    "difficultyWeight": 5
  },
  "Flaschenöffner": {
    "count": 6,
    "lastSeenTime": 1586874409884
  },
  "Roman": {
    "count": 2,
    "lastSeenTime": 1586839953759
  },
  "Nase": {
    "count": 6,
    "lastSeenTime": 1586916637868
  },
  "Regentropfen": {
    "count": 4,
    "lastSeenTime": 1586958141882,
    "difficulty": 1,
    "difficultyWeight": 10
  },
  "Statue": {
    "count": 7,
    "lastSeenTime": 1586920034941
  },
  "Studio": {
    "count": 5,
    "lastSeenTime": 1586889446349
  },
  "Portrait": {
    "count": 6,
    "lastSeenTime": 1586919399142,
    "difficulty": 1,
    "difficultyWeight": 32
  },
  "kabellos": {
    "count": 7,
    "lastSeenTime": 1586920034941,
    "difficulty": 1,
    "difficultyWeight": 6
  },
  "Zimmerpflanze": {
    "count": 2,
    "lastSeenTime": 1586919501543
  },
  "Prinzessin": {
    "count": 2,
    "lastSeenTime": 1586921008446
  },
  "Vollmond": {
    "count": 5,
    "lastSeenTime": 1586908395891,
    "difficulty": 0.9565217391304348,
    "difficultyWeight": 23
  },
  "Kupfer": {
    "count": 4,
    "lastSeenTime": 1586816578375
  },
  "Biologie": {
    "count": 3,
    "lastSeenTime": 1586954727368
  },
  "Papst": {
    "count": 5,
    "lastSeenTime": 1586840892251
  },
  "rutschig": {
    "count": 3,
    "lastSeenTime": 1586955113604
  },
  "Zyklop": {
    "count": 4,
    "lastSeenTime": 1586881056350,
    "publicGameCount": 1,
    "difficulty": 1,
    "difficultyWeight": 12
  },
  "Grotte": {
    "count": 4,
    "lastSeenTime": 1586915236139
  },
  "Stachelschwein": {
    "count": 3,
    "lastSeenTime": 1586839250104,
    "difficulty": 0.7142857142857143,
    "difficultyWeight": 7,
    "publicGameCount": 1
  },
  "Smaragd": {
    "count": 3,
    "lastSeenTime": 1586812763441
  },
  "Kriegsschiff": {
    "count": 5,
    "lastSeenTime": 1586858883401
  },
  "Dürre": {
    "count": 6,
    "lastSeenTime": 1586908276050
  },
  "Puma": {
    "count": 5,
    "lastSeenTime": 1586896290835
  },
  "Symphonie": {
    "count": 4,
    "lastSeenTime": 1586908315438
  },
  "ClickBait": {
    "count": 7,
    "lastSeenTime": 1586900856713,
    "difficulty": 0.8181818181818182,
    "difficultyWeight": 11
  },
  "Kuppel": {
    "count": 4,
    "lastSeenTime": 1586875034252
  },
  "applaudieren": {
    "count": 6,
    "lastSeenTime": 1586915821415
  },
  "Flügel": {
    "count": 5,
    "lastSeenTime": 1586907814926,
    "difficulty": 0.5833333333333333,
    "difficultyWeight": 12
  },
  "Frieden": {
    "count": 6,
    "lastSeenTime": 1586955224835,
    "difficulty": 0.6896551724137931,
    "difficultyWeight": 29
  },
  "Wort": {
    "count": 8,
    "lastSeenTime": 1586920456841
  },
  "Teller": {
    "count": 3,
    "lastSeenTime": 1586907682972,
    "difficulty": 0.4,
    "difficultyWeight": 10
  },
  "Teppich": {
    "count": 5,
    "lastSeenTime": 1586875186652,
    "difficulty": 1,
    "difficultyWeight": 8
  },
  "Schraube": {
    "count": 3,
    "lastSeenTime": 1586873383077
  },
  "Tag": {
    "count": 6,
    "lastSeenTime": 1586903370090
  },
  "Kopf": {
    "count": 13,
    "lastSeenTime": 1586920143399
  },
  "Ferien": {
    "count": 5,
    "lastSeenTime": 1586890156350
  },
  "Fischernetz": {
    "count": 2,
    "lastSeenTime": 1586782926490,
    "difficulty": 0.5555555555555556,
    "difficultyWeight": 9
  },
  "Wurst": {
    "count": 4,
    "lastSeenTime": 1586904367956,
    "difficulty": 1,
    "difficultyWeight": 3
  },
  "Triebwerk": {
    "count": 4,
    "lastSeenTime": 1586827603168
  },
  "Robbe": {
    "count": 7,
    "lastSeenTime": 1586921156587
  },
  "Wald": {
    "count": 3,
    "lastSeenTime": 1586712987795
  },
  "Netz": {
    "count": 8,
    "lastSeenTime": 1586920666083,
    "difficulty": 0.2857142857142857,
    "difficultyWeight": 7
  },
  "Fehler": {
    "count": 3,
    "lastSeenTime": 1586868805974
  },
  "Kastagnetten": {
    "count": 4,
    "lastSeenTime": 1586903148142
  },
  "Aladdin": {
    "count": 2,
    "lastSeenTime": 1586957874303,
    "difficulty": 1,
    "difficultyWeight": 13
  },
  "Scharfschütze": {
    "count": 3,
    "lastSeenTime": 1586826372203
  },
  "Korken": {
    "count": 8,
    "lastSeenTime": 1586916237094,
    "difficulty": 0.875,
    "difficultyWeight": 8
  },
  "Atlantis": {
    "count": 3,
    "lastSeenTime": 1586958572090
  },
  "Wüste": {
    "count": 5,
    "lastSeenTime": 1586859799165,
    "publicGameCount": 1,
    "difficulty": 0.8888888888888888,
    "difficultyWeight": 9
  },
  "Fernsehturm": {
    "count": 5,
    "lastSeenTime": 1586908382794
  },
  "Lautsprecher": {
    "count": 3,
    "lastSeenTime": 1586869993172,
    "difficulty": 0.8,
    "difficultyWeight": 5
  },
  "binden": {
    "count": 8,
    "lastSeenTime": 1586919773316,
    "difficulty": 1,
    "difficultyWeight": 2
  },
  "Pinguin": {
    "count": 4,
    "lastSeenTime": 1586901166652
  },
  "Gier": {
    "count": 6,
    "lastSeenTime": 1586902813258
  },
  "sterben": {
    "count": 6,
    "lastSeenTime": 1586915200929
  },
  "versagen": {
    "count": 6,
    "lastSeenTime": 1586956384816
  },
  "Unterschrift": {
    "count": 3,
    "lastSeenTime": 1586812458521
  },
  "Vuvuzela": {
    "count": 5,
    "lastSeenTime": 1586900861532
  },
  "Geografie": {
    "count": 6,
    "lastSeenTime": 1586896730722
  },
  "Oval": {
    "count": 1,
    "lastSeenTime": 1586618468051
  },
  "Marke": {
    "count": 5,
    "lastSeenTime": 1586914554202,
    "difficulty": 0.6666666666666666,
    "difficultyWeight": 15
  },
  "Observatorium": {
    "count": 2,
    "lastSeenTime": 1586909094851,
    "difficulty": 1,
    "difficultyWeight": 6
  },
  "Lesezeichen": {
    "count": 8,
    "lastSeenTime": 1586959455308,
    "difficulty": 1,
    "difficultyWeight": 7
  },
  "Pogo Stick": {
    "count": 6,
    "lastSeenTime": 1586902943577
  },
  "Zikade": {
    "count": 5,
    "lastSeenTime": 1586814522290
  },
  "Bulldozer": {
    "count": 5,
    "lastSeenTime": 1586955532027
  },
  "Gummi": {
    "count": 7,
    "lastSeenTime": 1586954678504
  },
  "Zuma": {
    "count": 5,
    "lastSeenTime": 1586908504643
  },
  "Saturn": {
    "count": 6,
    "lastSeenTime": 1586956352310,
    "difficulty": 1,
    "difficultyWeight": 7
  },
  "Tower Bridge": {
    "count": 7,
    "lastSeenTime": 1586914475739
  },
  "Zelt": {
    "count": 8,
    "lastSeenTime": 1586907327679,
    "difficulty": 1,
    "difficultyWeight": 23
  },
  "Darts": {
    "count": 3,
    "lastSeenTime": 1586650835510,
    "difficulty": 0.75,
    "difficultyWeight": 12
  },
  "Karaoke": {
    "count": 4,
    "lastSeenTime": 1586904384265
  },
  "Mitternacht": {
    "count": 5,
    "lastSeenTime": 1586955025294,
    "difficulty": 1,
    "difficultyWeight": 25
  },
  "Sechseck": {
    "count": 8,
    "lastSeenTime": 1586914157190
  },
  "Afrika": {
    "count": 2,
    "lastSeenTime": 1586895919376,
    "difficulty": 0.8,
    "difficultyWeight": 15,
    "publicGameCount": 1
  },
  "Arztkoffer": {
    "count": 5,
    "lastSeenTime": 1586889432199,
    "difficulty": 0.8,
    "difficultyWeight": 5
  },
  "Niere": {
    "count": 9,
    "lastSeenTime": 1586958385310
  },
  "Windrad": {
    "count": 4,
    "lastSeenTime": 1586869955906
  },
  "Baumkuchen": {
    "count": 4,
    "lastSeenTime": 1586900534749,
    "difficulty": 1,
    "difficultyWeight": 5
  },
  "Rasenmäher": {
    "count": 6,
    "lastSeenTime": 1586955691477
  },
  "Schiefer Turm von Pisa": {
    "count": 6,
    "lastSeenTime": 1586834077836
  },
  "Messer": {
    "count": 12,
    "lastSeenTime": 1586958636395
  },
  "Rechteck": {
    "count": 6,
    "lastSeenTime": 1586916744663
  },
  "behindert": {
    "count": 8,
    "lastSeenTime": 1586901239971,
    "publicGameCount": 1,
    "difficulty": 0.8333333333333334,
    "difficultyWeight": 6
  },
  "Hitze": {
    "count": 1,
    "lastSeenTime": 1586619376532,
    "difficulty": 0.6521739130434783,
    "difficultyWeight": 23
  },
  "König": {
    "count": 6,
    "lastSeenTime": 1586920604951,
    "publicGameCount": 1,
    "difficulty": 0.42857142857142855,
    "difficultyWeight": 7
  },
  "Akkordeon": {
    "count": 6,
    "lastSeenTime": 1586954577957
  },
  "Autogramm": {
    "count": 6,
    "lastSeenTime": 1586914716657,
    "publicGameCount": 1
  },
  "Dalmatiner": {
    "count": 7,
    "lastSeenTime": 1586955263523
  },
  "Fata Morgana": {
    "count": 4,
    "lastSeenTime": 1586958973713
  },
  "Meerschweinchen": {
    "count": 10,
    "lastSeenTime": 1586839691166,
    "publicGameCount": 1,
    "difficulty": 0.2857142857142857,
    "difficultyWeight": 7
  },
  "Pudel": {
    "count": 6,
    "lastSeenTime": 1586955310334
  },
  "Asien": {
    "count": 3,
    "lastSeenTime": 1586816712958,
    "difficulty": 0.625,
    "difficultyWeight": 8
  },
  "Haut": {
    "count": 5,
    "lastSeenTime": 1586900573287
  },
  "brünett": {
    "count": 4,
    "lastSeenTime": 1586920858569,
    "difficulty": 1,
    "difficultyWeight": 3
  },
  "Mohn": {
    "count": 3,
    "lastSeenTime": 1586889591248
  },
  "Aufnahme": {
    "count": 1,
    "lastSeenTime": 1586619614046
  },
  "Zuckerstange": {
    "count": 6,
    "lastSeenTime": 1586832545576
  },
  "Floß": {
    "count": 6,
    "lastSeenTime": 1586814801167
  },
  "Gips": {
    "count": 7,
    "lastSeenTime": 1586958719765
  },
  "rückgängig": {
    "count": 4,
    "lastSeenTime": 1586915456714,
    "difficulty": 1,
    "difficultyWeight": 9
  },
  "Schaltknüppel": {
    "count": 8,
    "lastSeenTime": 1586916288751,
    "difficulty": 0.6,
    "difficultyWeight": 5
  },
  "Poseidon": {
    "count": 10,
    "lastSeenTime": 1586956516194,
    "difficulty": 0.75,
    "difficultyWeight": 4
  },
  "Helm": {
    "count": 4,
    "lastSeenTime": 1586959106188
  },
  "Hütte": {
    "count": 4,
    "lastSeenTime": 1586832859836
  },
  "Tintenfisch": {
    "count": 3,
    "lastSeenTime": 1586901572201
  },
  "Skittles": {
    "count": 8,
    "lastSeenTime": 1586920722806
  },
  "Freddie Faulig": {
    "count": 5,
    "lastSeenTime": 1586921195186
  },
  "Bildhauer": {
    "count": 5,
    "lastSeenTime": 1586896885306
  },
  "Hund": {
    "count": 8,
    "lastSeenTime": 1586959142159,
    "difficulty": 1,
    "difficultyWeight": 3
  },
  "Dachfenster": {
    "count": 4,
    "lastSeenTime": 1586956197682,
    "difficulty": 1,
    "difficultyWeight": 1
  },
  "Postkarte": {
    "count": 3,
    "lastSeenTime": 1586896214025,
    "difficulty": 0.8545454545454545,
    "difficultyWeight": 55
  },
  "Großmutter": {
    "count": 5,
    "lastSeenTime": 1586903760979,
    "difficulty": 1,
    "difficultyWeight": 7
  },
  "Tochter": {
    "count": 4,
    "lastSeenTime": 1586920771641
  },
  "Motherboard": {
    "count": 2,
    "lastSeenTime": 1586724278124
  },
  "Panther": {
    "count": 2,
    "lastSeenTime": 1586908672662,
    "difficulty": 1,
    "difficultyWeight": 8
  },
  "Küstenwache": {
    "count": 4,
    "lastSeenTime": 1586813867768
  },
  "Chuck Norris": {
    "count": 5,
    "lastSeenTime": 1586914903455
  },
  "Hammer": {
    "count": 6,
    "lastSeenTime": 1586880115013,
    "publicGameCount": 1,
    "difficulty": 0.7272727272727273,
    "difficultyWeight": 11
  },
  "Apfelsaft": {
    "count": 7,
    "lastSeenTime": 1586903344033
  },
  "Kameramann": {
    "count": 5,
    "lastSeenTime": 1586956302197
  },
  "Gebäude": {
    "count": 4,
    "lastSeenTime": 1586816998752,
    "difficulty": 1,
    "difficultyWeight": 8
  },
  "Nickel": {
    "count": 4,
    "lastSeenTime": 1586879731749
  },
  "Schnuller": {
    "count": 5,
    "lastSeenTime": 1586891340621,
    "difficulty": 0.8928571428571429,
    "difficultyWeight": 28
  },
  "wählen": {
    "count": 6,
    "lastSeenTime": 1586916441133
  },
  "Pikachu": {
    "count": 5,
    "lastSeenTime": 1586958959180
  },
  "wiederholen": {
    "count": 7,
    "lastSeenTime": 1586874577908
  },
  "Amerika": {
    "count": 5,
    "lastSeenTime": 1586914615571,
    "difficulty": 1,
    "difficultyWeight": 3
  },
  "Zeh": {
    "count": 6,
    "lastSeenTime": 1586840668869,
    "difficulty": 0.5,
    "difficultyWeight": 8
  },
  "Skelett": {
    "count": 6,
    "lastSeenTime": 1586907814926
  },
  "Box": {
    "count": 6,
    "lastSeenTime": 1586956570180
  },
  "Hängebrücke": {
    "count": 7,
    "lastSeenTime": 1586903975283,
    "difficulty": 0.4,
    "difficultyWeight": 10
  },
  "Kätzchen": {
    "count": 2,
    "lastSeenTime": 1586891134706
  },
  "Bandscheibenvorfall": {
    "count": 9,
    "lastSeenTime": 1586900301076
  },
  "Pfeife": {
    "count": 2,
    "lastSeenTime": 1586900162376,
    "difficulty": 0.4,
    "difficultyWeight": 5
  },
  "Tauchen": {
    "count": 5,
    "lastSeenTime": 1586815355269
  },
  "Maschine": {
    "count": 8,
    "lastSeenTime": 1586915493381,
    "publicGameCount": 1,
    "difficulty": 0.75,
    "difficultyWeight": 4
  },
  "Sandalen": {
    "count": 7,
    "lastSeenTime": 1586896912568
  },
  "Sichel": {
    "count": 8,
    "lastSeenTime": 1586874931565
  },
  "Atombombe": {
    "count": 7,
    "lastSeenTime": 1586955457642
  },
  "Wachs": {
    "count": 3,
    "lastSeenTime": 1586958456673
  },
  "England": {
    "count": 7,
    "lastSeenTime": 1586957950175
  },
  "Temperatur": {
    "count": 5,
    "lastSeenTime": 1586958062580,
    "difficulty": 0.8888888888888888,
    "difficultyWeight": 9
  },
  "Gärtner": {
    "count": 6,
    "lastSeenTime": 1586907729754
  },
  "Nadelkissen": {
    "count": 5,
    "lastSeenTime": 1586902602080,
    "difficulty": 0.75,
    "difficultyWeight": 4
  },
  "Geschlecht": {
    "count": 3,
    "lastSeenTime": 1586806009929
  },
  "sitzen": {
    "count": 5,
    "lastSeenTime": 1586915924921
  },
  "Spaghettieis": {
    "count": 4,
    "lastSeenTime": 1586920375356
  },
  "Viertel": {
    "count": 5,
    "lastSeenTime": 1586901042640
  },
  "Nudel": {
    "count": 3,
    "lastSeenTime": 1586805458421,
    "difficulty": 1,
    "difficultyWeight": 8
  },
  "Maulwurf": {
    "count": 5,
    "lastSeenTime": 1586914157190,
    "difficulty": 0.6666666666666666,
    "difficultyWeight": 3
  },
  "Ferrari": {
    "count": 4,
    "lastSeenTime": 1586955823724,
    "difficulty": 0.95,
    "difficultyWeight": 20
  },
  "Cello": {
    "count": 2,
    "lastSeenTime": 1586711665898,
    "difficulty": 0.75,
    "difficultyWeight": 8
  },
  "Regisseur": {
    "count": 6,
    "lastSeenTime": 1586954953467
  },
  "Schuhkarton": {
    "count": 5,
    "lastSeenTime": 1586909286837
  },
  "Passwort": {
    "count": 6,
    "lastSeenTime": 1586958413299,
    "difficulty": 0.7142857142857143,
    "difficultyWeight": 7
  },
  "Markt": {
    "count": 5,
    "lastSeenTime": 1586826599494,
    "difficulty": 0.5,
    "difficultyWeight": 10
  },
  "Teleskop": {
    "count": 3,
    "lastSeenTime": 1586958485799,
    "difficulty": 0.875,
    "difficultyWeight": 8
  },
  "Töne": {
    "count": 2,
    "lastSeenTime": 1586868568396,
    "difficulty": 0.9285714285714286,
    "difficultyWeight": 14
  },
  "Schwarzes Loch": {
    "count": 5,
    "lastSeenTime": 1586955025293,
    "difficulty": 1,
    "difficultyWeight": 7
  },
  "Rumänien": {
    "count": 5,
    "lastSeenTime": 1586914501681
  },
  "Generator": {
    "count": 3,
    "lastSeenTime": 1586880812768,
    "difficulty": 1,
    "difficultyWeight": 9
  },
  "Leidenschaft": {
    "count": 5,
    "lastSeenTime": 1586958012186,
    "difficulty": 0.75,
    "difficultyWeight": 4
  },
  "Parkuhr": {
    "count": 4,
    "lastSeenTime": 1586958385310
  },
  "Universität": {
    "count": 4,
    "lastSeenTime": 1586814778569
  },
  "Rentier": {
    "count": 7,
    "lastSeenTime": 1586958525123
  },
  "Magie": {
    "count": 4,
    "lastSeenTime": 1586959242580
  },
  "Pumba": {
    "count": 5,
    "lastSeenTime": 1586959533143
  },
  "Zombie": {
    "count": 6,
    "lastSeenTime": 1586958547642,
    "publicGameCount": 1,
    "difficulty": 0.36363636363636365,
    "difficultyWeight": 11
  },
  "Maus": {
    "count": 4,
    "lastSeenTime": 1586920069685
  },
  "Darsteller": {
    "count": 8,
    "lastSeenTime": 1586959492934,
    "difficulty": 0.7857142857142857,
    "difficultyWeight": 14
  },
  "Norden": {
    "count": 3,
    "lastSeenTime": 1586817033855,
    "publicGameCount": 1
  },
  "Rasiermesser": {
    "count": 4,
    "lastSeenTime": 1586913973573
  },
  "Liane": {
    "count": 5,
    "lastSeenTime": 1586955457642,
    "difficulty": 0.8,
    "difficultyWeight": 5
  },
  "Ruhe": {
    "count": 3,
    "lastSeenTime": 1586813181137
  },
  "Kanone": {
    "count": 3,
    "lastSeenTime": 1586859826495
  },
  "Konversation": {
    "count": 4,
    "lastSeenTime": 1586826301547
  },
  "Windows": {
    "count": 4,
    "lastSeenTime": 1586955209647,
    "difficulty": 0.9852941176470589,
    "difficultyWeight": 68
  },
  "Villa": {
    "count": 3,
    "lastSeenTime": 1586919374724
  },
  "Papiertüte": {
    "count": 6,
    "lastSeenTime": 1586959304612
  },
  "rechts": {
    "count": 5,
    "lastSeenTime": 1586959608297
  },
  "Tornado": {
    "count": 4,
    "lastSeenTime": 1586874513513,
    "difficulty": 0.9354838709677419,
    "difficultyWeight": 31
  },
  "Kondenswasser": {
    "count": 6,
    "lastSeenTime": 1586890429362
  },
  "Sprühfarbe": {
    "count": 4,
    "lastSeenTime": 1586813950788,
    "difficulty": 1,
    "difficultyWeight": 10
  },
  "Ziel": {
    "count": 3,
    "lastSeenTime": 1586826637533,
    "difficulty": 1,
    "difficultyWeight": 10
  },
  "Sombrero": {
    "count": 7,
    "lastSeenTime": 1586958062580,
    "publicGameCount": 1
  },
  "Bushaltestelle": {
    "count": 3,
    "lastSeenTime": 1586783233936
  },
  "Purzelbaum": {
    "count": 3,
    "lastSeenTime": 1586915757073,
    "difficulty": 0.8571428571428571,
    "difficultyWeight": 7
  },
  "Haare": {
    "count": 1,
    "lastSeenTime": 1586634491034,
    "difficulty": 1,
    "difficultyWeight": 10
  },
  "Zypresse": {
    "count": 7,
    "lastSeenTime": 1586958809891
  },
  "Amsterdam": {
    "count": 4,
    "lastSeenTime": 1586920627218
  },
  "Makkaroni": {
    "count": 6,
    "lastSeenTime": 1586959830012
  },
  "taub": {
    "count": 4,
    "lastSeenTime": 1586880907839
  },
  "Bill Gates": {
    "count": 4,
    "lastSeenTime": 1586915129381
  },
  "provozieren": {
    "count": 5,
    "lastSeenTime": 1586896840770
  },
  "Raumanzug": {
    "count": 8,
    "lastSeenTime": 1586956019428,
    "difficulty": 1,
    "difficultyWeight": 82
  },
  "Punker": {
    "count": 4,
    "lastSeenTime": 1586959255115
  },
  "Thor": {
    "count": 4,
    "lastSeenTime": 1586880911950
  },
  "Globus": {
    "count": 5,
    "lastSeenTime": 1586889772751
  },
  "Tischplatte": {
    "count": 3,
    "lastSeenTime": 1586889906994
  },
  "hypnotisieren": {
    "count": 4,
    "lastSeenTime": 1586959401878
  },
  "Schnabeltier": {
    "count": 7,
    "lastSeenTime": 1586908611825
  },
  "Kreuzotter": {
    "count": 5,
    "lastSeenTime": 1586920957237,
    "publicGameCount": 1,
    "difficulty": 0.45454545454545453,
    "difficultyWeight": 11
  },
  "Sauna": {
    "count": 9,
    "lastSeenTime": 1586956363582
  },
  "Planke": {
    "count": 5,
    "lastSeenTime": 1586903323590
  },
  "bezaubernd": {
    "count": 5,
    "lastSeenTime": 1586889746364
  },
  "Nonne": {
    "count": 4,
    "lastSeenTime": 1586955958335
  },
  "Bücherei": {
    "count": 5,
    "lastSeenTime": 1586903231734
  },
  "Kakerlake": {
    "count": 3,
    "lastSeenTime": 1586827501149,
    "difficulty": 1,
    "difficultyWeight": 13
  },
  "Traubensaft": {
    "count": 5,
    "lastSeenTime": 1586881168379
  },
  "Wohnung": {
    "count": 6,
    "lastSeenTime": 1586839175430
  },
  "Fussel": {
    "count": 10,
    "lastSeenTime": 1586908054437
  },
  "Perücke": {
    "count": 6,
    "lastSeenTime": 1586955334605
  },
  "Querflöte": {
    "count": 5,
    "lastSeenTime": 1586921170820
  },
  "Vogelspinne": {
    "count": 7,
    "lastSeenTime": 1586908090329
  },
  "Kinn": {
    "count": 5,
    "lastSeenTime": 1586916063451,
    "difficulty": 0.8666666666666667,
    "difficultyWeight": 15
  },
  "Spinat": {
    "count": 6,
    "lastSeenTime": 1586959255115
  },
  "Glocke": {
    "count": 7,
    "lastSeenTime": 1586903888016
  },
  "Bäckerei": {
    "count": 5,
    "lastSeenTime": 1586921231641,
    "difficulty": 1,
    "difficultyWeight": 10
  },
  "Seehund": {
    "count": 5,
    "lastSeenTime": 1586959547988
  },
  "Konfetti": {
    "count": 4,
    "lastSeenTime": 1586919927137,
    "difficulty": 1,
    "difficultyWeight": 3
  },
  "Pfote": {
    "count": 4,
    "lastSeenTime": 1586869599214
  },
  "Nordkorea": {
    "count": 5,
    "lastSeenTime": 1586858711737
  },
  "Schlittschuh": {
    "count": 5,
    "lastSeenTime": 1586874687834
  },
  "Wolverine": {
    "count": 6,
    "lastSeenTime": 1586959429213
  },
  "Drucker": {
    "count": 2,
    "lastSeenTime": 1586713676145
  },
  "Elektriker": {
    "count": 6,
    "lastSeenTime": 1586959533143,
    "difficulty": 1,
    "difficultyWeight": 5
  },
  "Hausmeister": {
    "count": 5,
    "lastSeenTime": 1586958873120
  },
  "Schneeball": {
    "count": 9,
    "lastSeenTime": 1586920005337
  },
  "Zebra": {
    "count": 4,
    "lastSeenTime": 1586879903131,
    "difficulty": 0.25,
    "difficultyWeight": 8
  },
  "Ladebalken": {
    "count": 6,
    "lastSeenTime": 1586959218404,
    "difficulty": 0.875,
    "difficultyWeight": 8
  },
  "Blutegel": {
    "count": 10,
    "lastSeenTime": 1586920265574
  },
  "Feuerwehrauto": {
    "count": 7,
    "lastSeenTime": 1586873974505,
    "difficulty": 0.8974358974358975,
    "difficultyWeight": 39
  },
  "Architekt": {
    "count": 2,
    "lastSeenTime": 1586879685669,
    "difficulty": 1,
    "difficultyWeight": 4
  },
  "Kabine": {
    "count": 4,
    "lastSeenTime": 1586959850831
  },
  "Waffe": {
    "count": 3,
    "lastSeenTime": 1586895782227
  },
  "Heizungskessel": {
    "count": 5,
    "lastSeenTime": 1586913998392
  },
  "Recycling": {
    "count": 1,
    "lastSeenTime": 1586652085658
  },
  "Narwal": {
    "count": 9,
    "lastSeenTime": 1586920043786
  },
  "Würfel": {
    "count": 5,
    "lastSeenTime": 1586913696698
  },
  "tropisch": {
    "count": 4,
    "lastSeenTime": 1586908931787
  },
  "Bügeleisen": {
    "count": 3,
    "lastSeenTime": 1586955984837,
    "difficulty": 1,
    "difficultyWeight": 4
  },
  "Militär": {
    "count": 2,
    "lastSeenTime": 1586833977662
  },
  "Motor": {
    "count": 5,
    "lastSeenTime": 1586916288751
  },
  "Diagramm": {
    "count": 4,
    "lastSeenTime": 1586873676367
  },
  "Katamaran": {
    "count": 5,
    "lastSeenTime": 1586880408576
  },
  "Windschutzscheibe": {
    "count": 3,
    "lastSeenTime": 1586900587471
  },
  "Bob": {
    "count": 6,
    "lastSeenTime": 1586959279701,
    "publicGameCount": 1,
    "difficulty": 0.4166666666666667,
    "difficultyWeight": 12
  },
  "Scheibenwelt": {
    "count": 5,
    "lastSeenTime": 1586908191616
  },
  "Höhle": {
    "count": 7,
    "lastSeenTime": 1586956253026,
    "difficulty": 0.9090909090909091,
    "difficultyWeight": 11
  },
  "Werwolf": {
    "count": 5,
    "lastSeenTime": 1586916430600,
    "difficulty": 1,
    "difficultyWeight": 14
  },
  "Grapefruit": {
    "count": 3,
    "lastSeenTime": 1586833416769
  },
  "Waschmaschine": {
    "count": 7,
    "lastSeenTime": 1586909094851
  },
  "salto schlagen": {
    "count": 9,
    "lastSeenTime": 1586958496125
  },
  "Totenkopf": {
    "count": 4,
    "lastSeenTime": 1586900436732,
    "difficulty": 0.7391304347826086,
    "difficultyWeight": 23
  },
  "Gandalf": {
    "count": 6,
    "lastSeenTime": 1586916690790
  },
  "Spaten": {
    "count": 5,
    "lastSeenTime": 1586915788288,
    "difficulty": 0.7777777777777778,
    "difficultyWeight": 9
  },
  "Schlagzeug": {
    "count": 5,
    "lastSeenTime": 1586920771641
  },
  "Person": {
    "count": 7,
    "lastSeenTime": 1586954914968,
    "difficulty": 0.7586206896551724,
    "difficultyWeight": 29
  },
  "Büro": {
    "count": 5,
    "lastSeenTime": 1586838842778
  },
  "Ameise": {
    "count": 4,
    "lastSeenTime": 1586959314784,
    "difficulty": 0.7,
    "difficultyWeight": 10
  },
  "Nilpferd": {
    "count": 8,
    "lastSeenTime": 1586958821795,
    "difficulty": 0.9565217391304348,
    "difficultyWeight": 46
  },
  "Krabbe": {
    "count": 5,
    "lastSeenTime": 1586954798444
  },
  "Granatapfel": {
    "count": 8,
    "lastSeenTime": 1586919460885
  },
  "Tumor": {
    "count": 7,
    "lastSeenTime": 1586908480352
  },
  "Pi": {
    "count": 6,
    "lastSeenTime": 1586915129381
  },
  "Pferd": {
    "count": 2,
    "lastSeenTime": 1586858562011
  },
  "Textmarker": {
    "count": 2,
    "lastSeenTime": 1586709415728,
    "difficulty": 1,
    "difficultyWeight": 12
  },
  "Student": {
    "count": 4,
    "lastSeenTime": 1586955542167
  },
  "schwanger": {
    "count": 5,
    "lastSeenTime": 1586955701781,
    "difficulty": 1,
    "difficultyWeight": 6
  },
  "Gymnastik": {
    "count": 6,
    "lastSeenTime": 1586955195086,
    "difficulty": 0.7142857142857143,
    "difficultyWeight": 7
  },
  "Fackel": {
    "count": 4,
    "lastSeenTime": 1586859593426
  },
  "Essstäbchen": {
    "count": 4,
    "lastSeenTime": 1586897115648,
    "difficulty": 1,
    "difficultyWeight": 3
  },
  "Wellensittich": {
    "count": 7,
    "lastSeenTime": 1586956070276
  },
  "Wildnis": {
    "count": 3,
    "lastSeenTime": 1586889432199
  },
  "Strumpfhose": {
    "count": 5,
    "lastSeenTime": 1586958934411,
    "difficulty": 0.75,
    "difficultyWeight": 12
  },
  "Laktoseintoleranz": {
    "count": 4,
    "lastSeenTime": 1586879349367
  },
  "Knallfrosch": {
    "count": 3,
    "lastSeenTime": 1586713746274
  },
  "suchen": {
    "count": 4,
    "lastSeenTime": 1586957960523
  },
  "Eichel": {
    "count": 8,
    "lastSeenTime": 1586903939910,
    "difficulty": 0.9166666666666666,
    "difficultyWeight": 48
  },
  "Hose": {
    "count": 8,
    "lastSeenTime": 1586959178759,
    "difficulty": 0.45,
    "difficultyWeight": 40
  },
  "Brokkoli": {
    "count": 4,
    "lastSeenTime": 1586955063993
  },
  "Blitz": {
    "count": 8,
    "lastSeenTime": 1586959522880,
    "difficulty": 1,
    "difficultyWeight": 5
  },
  "Excalibur": {
    "count": 3,
    "lastSeenTime": 1586908207834,
    "difficulty": 0.75,
    "difficultyWeight": 8
  },
  "Creeper": {
    "count": 6,
    "lastSeenTime": 1586833067061,
    "difficulty": 0.7272727272727273,
    "difficultyWeight": 11
  },
  "Röstaroma": {
    "count": 7,
    "lastSeenTime": 1586958335993
  },
  "Sahne": {
    "count": 7,
    "lastSeenTime": 1586959557107
  },
  "abstrakt": {
    "count": 7,
    "lastSeenTime": 1586915788288
  },
  "Bagger": {
    "count": 8,
    "lastSeenTime": 1586916579758
  },
  "Dorf": {
    "count": 5,
    "lastSeenTime": 1586955419506,
    "difficulty": 0.92,
    "difficultyWeight": 25
  },
  "Mineralwasser": {
    "count": 6,
    "lastSeenTime": 1586901647547,
    "difficulty": 0.9166666666666666,
    "difficultyWeight": 12
  },
  "Regenwald": {
    "count": 3,
    "lastSeenTime": 1586860432423
  },
  "Frühlingsrolle": {
    "count": 5,
    "lastSeenTime": 1586903034208
  },
  "Abonnement": {
    "count": 5,
    "lastSeenTime": 1586900318835,
    "difficulty": 0.375,
    "difficultyWeight": 8
  },
  "Stubenhocker": {
    "count": 6,
    "lastSeenTime": 1586895895138
  },
  "berühren": {
    "count": 4,
    "lastSeenTime": 1586958360375
  },
  "Schraubenschlüssel": {
    "count": 1,
    "lastSeenTime": 1586707602434
  },
  "Kurzschluss": {
    "count": 4,
    "lastSeenTime": 1586900212280
  },
  "Jacke": {
    "count": 5,
    "lastSeenTime": 1586955098414
  },
  "Samen": {
    "count": 3,
    "lastSeenTime": 1586868411733
  },
  "Narr": {
    "count": 2,
    "lastSeenTime": 1586895651423
  },
  "Kronkorken": {
    "count": 3,
    "lastSeenTime": 1586869743912
  },
  "Blüte": {
    "count": 4,
    "lastSeenTime": 1586827737429
  },
  "Ladegerät": {
    "count": 5,
    "lastSeenTime": 1586914932818
  },
  "weinen": {
    "count": 6,
    "lastSeenTime": 1586956070276,
    "difficulty": 1,
    "difficultyWeight": 3
  },
  "Sau": {
    "count": 7,
    "lastSeenTime": 1586920143399
  },
  "obdachlos": {
    "count": 5,
    "lastSeenTime": 1586879940075
  },
  "Einbahnstraße": {
    "count": 3,
    "lastSeenTime": 1586956384816
  },
  "Stangenbrot": {
    "count": 1,
    "lastSeenTime": 1586708591047
  },
  "Odysseus": {
    "count": 4,
    "lastSeenTime": 1586959492934
  },
  "London Eye": {
    "count": 4,
    "lastSeenTime": 1586955620355,
    "difficulty": 0.75,
    "difficultyWeight": 4
  },
  "Hundehütte": {
    "count": 4,
    "lastSeenTime": 1586913545993
  },
  "Eisberg": {
    "count": 5,
    "lastSeenTime": 1586915639271,
    "difficulty": 0.7272727272727273,
    "difficultyWeight": 11
  },
  "Evolution": {
    "count": 4,
    "lastSeenTime": 1586815367180,
    "difficulty": 0.7857142857142857,
    "difficultyWeight": 14
  },
  "Beziehung": {
    "count": 9,
    "lastSeenTime": 1586900559143
  },
  "Koffer": {
    "count": 5,
    "lastSeenTime": 1586954698858,
    "difficulty": 0.8,
    "difficultyWeight": 30
  },
  "Meerenge": {
    "count": 8,
    "lastSeenTime": 1586908583500,
    "difficulty": 1,
    "difficultyWeight": 8
  },
  "Posaune": {
    "count": 4,
    "lastSeenTime": 1586806917088
  },
  "Sturm": {
    "count": 2,
    "lastSeenTime": 1586920133284
  },
  "Lächeln": {
    "count": 1,
    "lastSeenTime": 1586708983645
  },
  "Star Wars": {
    "count": 4,
    "lastSeenTime": 1586919917033,
    "difficulty": 0.42857142857142855,
    "difficultyWeight": 7
  },
  "Halskette": {
    "count": 5,
    "lastSeenTime": 1586916655791
  },
  "Spore": {
    "count": 4,
    "lastSeenTime": 1586890203054,
    "difficulty": 1,
    "difficultyWeight": 7
  },
  "Ass": {
    "count": 6,
    "lastSeenTime": 1586916655791
  },
  "Böller": {
    "count": 3,
    "lastSeenTime": 1586919511658,
    "difficulty": 1,
    "difficultyWeight": 6
  },
  "Flohmarkt": {
    "count": 5,
    "lastSeenTime": 1586957837428
  },
  "Einkaufszentrum": {
    "count": 4,
    "lastSeenTime": 1586909133282,
    "difficulty": 1,
    "difficultyWeight": 7
  },
  "Locher": {
    "count": 3,
    "lastSeenTime": 1586889772751
  },
  "Fliege": {
    "count": 4,
    "lastSeenTime": 1586827275317,
    "difficulty": 0.75,
    "difficultyWeight": 8
  },
  "Kastanie": {
    "count": 5,
    "lastSeenTime": 1586921057562
  },
  "Kernkraftwerk": {
    "count": 5,
    "lastSeenTime": 1586915354278
  },
  "Wäschespinne": {
    "count": 4,
    "lastSeenTime": 1586907778315
  },
  "Strohhalm": {
    "count": 6,
    "lastSeenTime": 1586901264460
  },
  "Räuber": {
    "count": 5,
    "lastSeenTime": 1586958596468
  },
  "Mönch": {
    "count": 5,
    "lastSeenTime": 1586868078333
  },
  "Deutschland": {
    "count": 7,
    "lastSeenTime": 1586921231641,
    "difficulty": 0.75,
    "difficultyWeight": 8
  },
  "Schleuder": {
    "count": 3,
    "lastSeenTime": 1586903624140,
    "difficulty": 0.5625,
    "difficultyWeight": 32
  },
  "Terminator": {
    "count": 5,
    "lastSeenTime": 1586959090713
  },
  "Beatbox": {
    "count": 4,
    "lastSeenTime": 1586956273734,
    "difficulty": 0.8333333333333334,
    "difficultyWeight": 12
  },
  "Miss Piggy": {
    "count": 6,
    "lastSeenTime": 1586920153587
  },
  "Facebook": {
    "count": 6,
    "lastSeenTime": 1586816946669
  },
  "Schreibschrift": {
    "count": 3,
    "lastSeenTime": 1586920326688,
    "publicGameCount": 1,
    "difficulty": 0.8571428571428571,
    "difficultyWeight": 7
  },
  "Katalog": {
    "count": 3,
    "lastSeenTime": 1586921008446
  },
  "Gasmaske": {
    "count": 4,
    "lastSeenTime": 1586901500696
  },
  "Arbeiter": {
    "count": 2,
    "lastSeenTime": 1586913698511
  },
  "Mädchen": {
    "count": 7,
    "lastSeenTime": 1586958510664
  },
  "Restaurant": {
    "count": 3,
    "lastSeenTime": 1586954953467
  },
  "Achsel": {
    "count": 7,
    "lastSeenTime": 1586955923939,
    "difficulty": 0.5,
    "difficultyWeight": 6
  },
  "Burrito": {
    "count": 6,
    "lastSeenTime": 1586908686991,
    "difficulty": 0.9411764705882353,
    "difficultyWeight": 34
  },
  "Lederhose": {
    "count": 2,
    "lastSeenTime": 1586891124151
  },
  "UFO": {
    "count": 2,
    "lastSeenTime": 1586867998752,
    "difficulty": 0.5,
    "difficultyWeight": 14
  },
  "See": {
    "count": 9,
    "lastSeenTime": 1586955310334,
    "difficulty": 1,
    "difficultyWeight": 13
  },
  "Vogelbeere": {
    "count": 5,
    "lastSeenTime": 1586903221605
  },
  "Gipfelkreuz": {
    "count": 6,
    "lastSeenTime": 1586954989864
  },
  "Hai": {
    "count": 6,
    "lastSeenTime": 1586954808666
  },
  "grinsen": {
    "count": 6,
    "lastSeenTime": 1586916288751,
    "difficulty": 0.3333333333333333,
    "difficultyWeight": 6
  },
  "Youtube": {
    "count": 5,
    "lastSeenTime": 1586954641977
  },
  "Aufzug": {
    "count": 4,
    "lastSeenTime": 1586921156587
  },
  "Internet": {
    "count": 3,
    "lastSeenTime": 1586869993172
  },
  "Apfelkern": {
    "count": 2,
    "lastSeenTime": 1586727732259
  },
  "Leiterbahn": {
    "count": 2,
    "lastSeenTime": 1586956504597,
    "difficulty": 1,
    "difficultyWeight": 6
  },
  "Stephen Hawking": {
    "count": 3,
    "lastSeenTime": 1586907532918
  },
  "Topf": {
    "count": 7,
    "lastSeenTime": 1586902587809
  },
  "Degen": {
    "count": 5,
    "lastSeenTime": 1586897394573
  },
  "Zahnbürste": {
    "count": 12,
    "lastSeenTime": 1586914806527,
    "difficulty": 0.8518518518518519,
    "difficultyWeight": 27
  },
  "knien": {
    "count": 3,
    "lastSeenTime": 1586955195086
  },
  "Geschäft": {
    "count": 5,
    "lastSeenTime": 1586959804291,
    "difficulty": 0.8571428571428571,
    "difficultyWeight": 7
  },
  "Marge Simpson": {
    "count": 8,
    "lastSeenTime": 1586956516194,
    "difficulty": 0.5,
    "difficultyWeight": 6
  },
  "Stoppschild": {
    "count": 6,
    "lastSeenTime": 1586901606974
  },
  "CO2": {
    "count": 3,
    "lastSeenTime": 1586957989157,
    "difficulty": 1,
    "difficultyWeight": 3
  },
  "Schlacht": {
    "count": 5,
    "lastSeenTime": 1586957863554
  },
  "Wurzel": {
    "count": 2,
    "lastSeenTime": 1586903877862
  },
  "Jazz": {
    "count": 3,
    "lastSeenTime": 1586901284755
  },
  "Margarine": {
    "count": 4,
    "lastSeenTime": 1586813965844
  },
  "asymmetrisch": {
    "count": 4,
    "lastSeenTime": 1586891113896
  },
  "Natur": {
    "count": 3,
    "lastSeenTime": 1586839235965,
    "difficulty": 1,
    "difficultyWeight": 7
  },
  "Kamin": {
    "count": 3,
    "lastSeenTime": 1586903750797
  },
  "Vulkan": {
    "count": 3,
    "lastSeenTime": 1586815464569,
    "difficulty": 0.75,
    "difficultyWeight": 52
  },
  "read": {
    "count": 1,
    "lastSeenTime": 1586711191040,
    "difficulty": 0.9333333333333333,
    "difficultyWeight": 15
  },
  "verschwinden": {
    "count": 8,
    "lastSeenTime": 1586904051087
  },
  "skull": {
    "count": 1,
    "lastSeenTime": 1586711214119,
    "difficulty": 0.8666666666666667,
    "difficultyWeight": 15
  },
  "Vermieter": {
    "count": 3,
    "lastSeenTime": 1586838964600
  },
  "tennis": {
    "count": 1,
    "lastSeenTime": 1586711276270,
    "difficulty": 0.6470588235294117,
    "difficultyWeight": 17
  },
  "axe": {
    "count": 1,
    "lastSeenTime": 1586711304408,
    "difficulty": 0.8823529411764706,
    "difficultyWeight": 17
  },
  "book": {
    "count": 1,
    "lastSeenTime": 1586711308520
  },
  "glass": {
    "count": 1,
    "lastSeenTime": 1586711308520
  },
  "Sprungturm": {
    "count": 6,
    "lastSeenTime": 1586955074091
  },
  "Professor": {
    "count": 5,
    "lastSeenTime": 1586880438207
  },
  "Spion": {
    "count": 2,
    "lastSeenTime": 1586859922496
  },
  "Säure": {
    "count": 5,
    "lastSeenTime": 1586916717644
  },
  "Sonntag": {
    "count": 2,
    "lastSeenTime": 1586915048330
  },
  "Kendama": {
    "count": 6,
    "lastSeenTime": 1586921057562
  },
  "Montag": {
    "count": 6,
    "lastSeenTime": 1586954688633
  },
  "Brusthaare": {
    "count": 7,
    "lastSeenTime": 1586879731749
  },
  "Brause": {
    "count": 6,
    "lastSeenTime": 1586916600438
  },
  "verletzt": {
    "count": 7,
    "lastSeenTime": 1586919511658,
    "difficulty": 1,
    "difficultyWeight": 17
  },
  "Specht": {
    "count": 5,
    "lastSeenTime": 1586827905301
  },
  "Nuss": {
    "count": 2,
    "lastSeenTime": 1586840446597,
    "difficulty": 1,
    "difficultyWeight": 7
  },
  "Skype": {
    "count": 4,
    "lastSeenTime": 1586958611285
  },
  "Rasensprenger": {
    "count": 6,
    "lastSeenTime": 1586914526456,
    "difficulty": 0.875,
    "difficultyWeight": 8
  },
  "Indianer": {
    "count": 2,
    "lastSeenTime": 1586718330257
  },
  "Geschenk": {
    "count": 6,
    "lastSeenTime": 1586903400691,
    "difficulty": 0.8181818181818182,
    "difficultyWeight": 11
  },
  "Einstein": {
    "count": 1,
    "lastSeenTime": 1586712023064,
    "difficulty": 1,
    "difficultyWeight": 5
  },
  "Kennzeichen": {
    "count": 3,
    "lastSeenTime": 1586955644812,
    "difficulty": 1,
    "difficultyWeight": 13
  },
  "Baby": {
    "count": 4,
    "lastSeenTime": 1586920843703
  },
  "Pasta": {
    "count": 2,
    "lastSeenTime": 1586839545523
  },
  "Pistazie": {
    "count": 4,
    "lastSeenTime": 1586915741689
  },
  "niesen": {
    "count": 5,
    "lastSeenTime": 1586903059681,
    "difficulty": 0.9807692307692307,
    "difficultyWeight": 52
  },
  "Nadel": {
    "count": 3,
    "lastSeenTime": 1586921022722
  },
  "Biotonne": {
    "count": 2,
    "lastSeenTime": 1586896127512
  },
  "Pinzette": {
    "count": 4,
    "lastSeenTime": 1586869906740
  },
  "Notizbuch": {
    "count": 3,
    "lastSeenTime": 1586833628220
  },
  "krank": {
    "count": 6,
    "lastSeenTime": 1586908611825
  },
  "Klavier": {
    "count": 4,
    "lastSeenTime": 1586908707244
  },
  "Salat": {
    "count": 3,
    "lastSeenTime": 1586813485208
  },
  "Geologe": {
    "count": 6,
    "lastSeenTime": 1586919640209,
    "difficulty": 1,
    "difficultyWeight": 6
  },
  "Trampolin": {
    "count": 2,
    "lastSeenTime": 1586813624926,
    "difficulty": 0.8,
    "difficultyWeight": 5
  },
  "Iron Giant": {
    "count": 5,
    "lastSeenTime": 1586890368170
  },
  "sinken": {
    "count": 4,
    "lastSeenTime": 1586908079268,
    "difficulty": 0.625,
    "difficultyWeight": 8
  },
  "Schaufel": {
    "count": 3,
    "lastSeenTime": 1586958062580
  },
  "Soße": {
    "count": 5,
    "lastSeenTime": 1586869017466
  },
  "Stöpsel": {
    "count": 3,
    "lastSeenTime": 1586827250323
  },
  "Operation": {
    "count": 2,
    "lastSeenTime": 1586958755749
  },
  "Litfaßsäule": {
    "count": 5,
    "lastSeenTime": 1586812468912
  },
  "Soldat": {
    "count": 6,
    "lastSeenTime": 1586903634298,
    "difficulty": 0.7272727272727273,
    "difficultyWeight": 11
  },
  "Karneval": {
    "count": 7,
    "lastSeenTime": 1586908518895
  },
  "Küche": {
    "count": 7,
    "lastSeenTime": 1586916063451,
    "difficulty": 1,
    "difficultyWeight": 9
  },
  "Olive": {
    "count": 2,
    "lastSeenTime": 1586900473463,
    "difficulty": 1,
    "difficultyWeight": 14
  },
  "Regenschirm": {
    "count": 2,
    "lastSeenTime": 1586901225801,
    "difficulty": 0.631578947368421,
    "difficultyWeight": 19
  },
  "Rückgrat": {
    "count": 5,
    "lastSeenTime": 1586897383053
  },
  "Eisverkäufer": {
    "count": 2,
    "lastSeenTime": 1586900763534,
    "difficulty": 0.6818181818181818,
    "difficultyWeight": 22
  },
  "Sicherheitsnadel": {
    "count": 1,
    "lastSeenTime": 1586713732056
  },
  "Merkur": {
    "count": 6,
    "lastSeenTime": 1586954641977
  },
  "Buntstift": {
    "count": 8,
    "lastSeenTime": 1586958200965,
    "difficulty": 1,
    "difficultyWeight": 9
  },
  "Achteck": {
    "count": 5,
    "lastSeenTime": 1586900212280,
    "publicGameCount": 1,
    "difficulty": 0.5833333333333333,
    "difficultyWeight": 12
  },
  "Gurke": {
    "count": 9,
    "lastSeenTime": 1586955037756
  },
  "Zitrone": {
    "count": 3,
    "lastSeenTime": 1586826941060
  },
  "Nacken": {
    "count": 3,
    "lastSeenTime": 1586913816802,
    "difficulty": 0.7,
    "difficultyWeight": 10
  },
  "Film": {
    "count": 6,
    "lastSeenTime": 1586959622590
  },
  "Kind": {
    "count": 6,
    "lastSeenTime": 1586920771641
  },
  "weiß": {
    "count": 5,
    "lastSeenTime": 1586907500591,
    "difficulty": 0.5,
    "difficultyWeight": 6
  },
  "Totem": {
    "count": 4,
    "lastSeenTime": 1586913962978,
    "difficulty": 0.3333333333333333,
    "difficultyWeight": 3
  },
  "trainieren": {
    "count": 10,
    "lastSeenTime": 1586958200965
  },
  "Sturz": {
    "count": 4,
    "lastSeenTime": 1586919522211
  },
  "Cheerleader": {
    "count": 5,
    "lastSeenTime": 1586916277356
  },
  "Adidas": {
    "count": 4,
    "lastSeenTime": 1586869335019,
    "publicGameCount": 1
  },
  "Fritten": {
    "count": 3,
    "lastSeenTime": 1586955344691
  },
  "kiffen": {
    "count": 7,
    "lastSeenTime": 1586896912568
  },
  "Schweden": {
    "count": 3,
    "lastSeenTime": 1586921059401
  },
  "Gelee": {
    "count": 6,
    "lastSeenTime": 1586915478982
  },
  "Gentleman": {
    "count": 2,
    "lastSeenTime": 1586869144958
  },
  "Kobra": {
    "count": 6,
    "lastSeenTime": 1586913497512,
    "difficulty": 0.5517241379310345,
    "difficultyWeight": 29
  },
  "BMW": {
    "count": 3,
    "lastSeenTime": 1586908017512,
    "difficulty": 1,
    "difficultyWeight": 14
  },
  "Feld": {
    "count": 6,
    "lastSeenTime": 1586955569169
  },
  "Athlet": {
    "count": 3,
    "lastSeenTime": 1586859853848
  },
  "Würfelqualle": {
    "count": 4,
    "lastSeenTime": 1586901839909
  },
  "Gekritzel": {
    "count": 5,
    "lastSeenTime": 1586920954248
  },
  "Birne": {
    "count": 7,
    "lastSeenTime": 1586908763208
  },
  "Wespe": {
    "count": 4,
    "lastSeenTime": 1586908432802,
    "difficulty": 0.75,
    "difficultyWeight": 24
  },
  "Karussell": {
    "count": 5,
    "lastSeenTime": 1586909155516
  },
  "Schwein": {
    "count": 4,
    "lastSeenTime": 1586921043126
  },
  "Hochzeit": {
    "count": 5,
    "lastSeenTime": 1586958299322
  },
  "Äffchen": {
    "count": 2,
    "lastSeenTime": 1586827841092
  },
  "Inlineskates": {
    "count": 4,
    "lastSeenTime": 1586919487290,
    "difficulty": 0.9,
    "difficultyWeight": 10
  },
  "Wetterfrosch": {
    "count": 4,
    "lastSeenTime": 1586904262976
  },
  "Taxifahrer": {
    "count": 8,
    "lastSeenTime": 1586955999033
  },
  "Einschränkung": {
    "count": 5,
    "lastSeenTime": 1586959522880
  },
  "verwirrt": {
    "count": 2,
    "lastSeenTime": 1586900387786
  },
  "flüssig": {
    "count": 2,
    "lastSeenTime": 1586875135071
  },
  "Schauspieler": {
    "count": 6,
    "lastSeenTime": 1586900278764
  },
  "U-Boot": {
    "count": 3,
    "lastSeenTime": 1586834009704
  },
  "Mandel": {
    "count": 1,
    "lastSeenTime": 1586723339566
  },
  "Dynamit": {
    "count": 6,
    "lastSeenTime": 1586908165354
  },
  "Weihnachtsmann": {
    "count": 7,
    "lastSeenTime": 1586915609902,
    "difficulty": 0.25,
    "difficultyWeight": 8
  },
  "Mantel": {
    "count": 6,
    "lastSeenTime": 1586919610641,
    "difficulty": 0.125,
    "difficultyWeight": 8
  },
  "Zahnstein": {
    "count": 3,
    "lastSeenTime": 1586899955403
  },
  "tauchen": {
    "count": 3,
    "lastSeenTime": 1586915089601
  },
  "Westen": {
    "count": 7,
    "lastSeenTime": 1586954939261,
    "difficulty": 0.55,
    "difficultyWeight": 20
  },
  "Zwischendecke": {
    "count": 6,
    "lastSeenTime": 1586959429213
  },
  "Pepsi": {
    "count": 3,
    "lastSeenTime": 1586908646355
  },
  "Busfahrer": {
    "count": 4,
    "lastSeenTime": 1586858858134
  },
  "Ringelschwanz": {
    "count": 3,
    "lastSeenTime": 1586896203705
  },
  "aufwerten": {
    "count": 5,
    "lastSeenTime": 1586956130879
  },
  "Efeu": {
    "count": 7,
    "lastSeenTime": 1586921256075
  },
  "Knoblauch": {
    "count": 3,
    "lastSeenTime": 1586916717644
  },
  "Meteorologe": {
    "count": 5,
    "lastSeenTime": 1586818368744,
    "difficulty": 1,
    "difficultyWeight": 6
  },
  "Nachtisch": {
    "count": 5,
    "lastSeenTime": 1586958037635
  },
  "Warze": {
    "count": 3,
    "lastSeenTime": 1586904186629,
    "difficulty": 1,
    "difficultyWeight": 7
  },
  "Pu der Bär": {
    "count": 2,
    "lastSeenTime": 1586726449700
  },
  "Bauarbeiter": {
    "count": 7,
    "lastSeenTime": 1586956384816
  },
  "Programmierer": {
    "count": 4,
    "lastSeenTime": 1586915985961
  },
  "Angebot": {
    "count": 5,
    "lastSeenTime": 1586903048902
  },
  "Stau": {
    "count": 4,
    "lastSeenTime": 1586916579758
  },
  "Antivirus": {
    "count": 2,
    "lastSeenTime": 1586805694894
  },
  "Anzeige": {
    "count": 8,
    "lastSeenTime": 1586916141879
  },
  "Museum": {
    "count": 2,
    "lastSeenTime": 1586914491372,
    "difficulty": 0.8809523809523809,
    "difficultyWeight": 42
  },
  "Anakonda": {
    "count": 6,
    "lastSeenTime": 1586954763962
  },
  "Suezkanal": {
    "count": 4,
    "lastSeenTime": 1586868821062
  },
  "Terrarium": {
    "count": 3,
    "lastSeenTime": 1586840867744
  },
  "Äquator": {
    "count": 3,
    "lastSeenTime": 1586919413397
  },
  "Zecke": {
    "count": 3,
    "lastSeenTime": 1586889746630
  },
  "Haarspray": {
    "count": 4,
    "lastSeenTime": 1586916252219
  },
  "Golf": {
    "count": 2,
    "lastSeenTime": 1586959232575
  },
  "Island": {
    "count": 4,
    "lastSeenTime": 1586956215415
  },
  "Simon and Garfunkel": {
    "count": 3,
    "lastSeenTime": 1586832474075
  },
  "Zebrastreifen": {
    "count": 4,
    "lastSeenTime": 1586880491865
  },
  "spülen": {
    "count": 4,
    "lastSeenTime": 1586955249129,
    "difficulty": 0.625,
    "difficultyWeight": 8
  },
  "Mundharmonika": {
    "count": 6,
    "lastSeenTime": 1586909211843
  },
  "fliegen": {
    "count": 3,
    "lastSeenTime": 1586921283836,
    "difficulty": 0.8571428571428571,
    "difficultyWeight": 14
  },
  "begrüßen": {
    "count": 4,
    "lastSeenTime": 1586920029620,
    "difficulty": 1,
    "difficultyWeight": 6
  },
  "Vanille": {
    "count": 3,
    "lastSeenTime": 1586812607253
  },
  "Emu": {
    "count": 3,
    "lastSeenTime": 1586873739947
  },
  "Gangster": {
    "count": 4,
    "lastSeenTime": 1586815888370
  },
  "Krümelmonster": {
    "count": 5,
    "lastSeenTime": 1586868628439,
    "difficulty": 0.8666666666666667,
    "difficultyWeight": 15
  },
  "Skilift": {
    "count": 5,
    "lastSeenTime": 1586913545993
  },
  "Verlierer": {
    "count": 5,
    "lastSeenTime": 1586869966082
  },
  "Gumball": {
    "count": 4,
    "lastSeenTime": 1586958668588
  },
  "Uhr": {
    "count": 4,
    "lastSeenTime": 1586958907689,
    "difficulty": 0.33333333333333337,
    "difficultyWeight": 9
  },
  "Kebab": {
    "count": 6,
    "lastSeenTime": 1586919750850,
    "difficulty": 1,
    "difficultyWeight": 4
  },
  "Tiger": {
    "count": 4,
    "lastSeenTime": 1586907863533
  },
  "Journalist": {
    "count": 4,
    "lastSeenTime": 1586896644364
  },
  "Finn": {
    "count": 2,
    "lastSeenTime": 1586914991061
  },
  "Salami": {
    "count": 6,
    "lastSeenTime": 1586907900355
  },
  "Pendel": {
    "count": 6,
    "lastSeenTime": 1586957974879
  },
  "Wettervorhersage": {
    "count": 6,
    "lastSeenTime": 1586919664805
  },
  "Stempel": {
    "count": 3,
    "lastSeenTime": 1586879283630
  },
  "Bettdecke": {
    "count": 6,
    "lastSeenTime": 1586916168548,
    "difficulty": 0.8461538461538461,
    "difficultyWeight": 13
  },
  "Halskrause": {
    "count": 5,
    "lastSeenTime": 1586904384265
  },
  "Albatros": {
    "count": 4,
    "lastSeenTime": 1586956552197
  },
  "Bücherwurm": {
    "count": 1,
    "lastSeenTime": 1586727461944
  },
  "Mikroskop": {
    "count": 5,
    "lastSeenTime": 1586958012186
  },
  "Läufer": {
    "count": 6,
    "lastSeenTime": 1586959767498
  },
  "Creme": {
    "count": 6,
    "lastSeenTime": 1586901152323
  },
  "Libelle": {
    "count": 4,
    "lastSeenTime": 1586839691166,
    "difficulty": 1,
    "difficultyWeight": 8
  },
  "Segelboot": {
    "count": 3,
    "lastSeenTime": 1586840135568,
    "difficulty": 1,
    "difficultyWeight": 39
  },
  "Hängebauchschwein": {
    "count": 4,
    "lastSeenTime": 1586908395891
  },
  "Radiergummi": {
    "count": 5,
    "lastSeenTime": 1586916252219
  },
  "Bumerang": {
    "count": 3,
    "lastSeenTime": 1586873278128
  },
  "schwindelig": {
    "count": 6,
    "lastSeenTime": 1586920108914
  },
  "Ohrwurm": {
    "count": 2,
    "lastSeenTime": 1586814854011
  },
  "Schwamm": {
    "count": 4,
    "lastSeenTime": 1586920926700
  },
  "flüstern": {
    "count": 3,
    "lastSeenTime": 1586903447898
  },
  "Modedesigner": {
    "count": 5,
    "lastSeenTime": 1586958239433
  },
  "Hip Hop": {
    "count": 2,
    "lastSeenTime": 1586813888653,
    "difficulty": 0.8,
    "difficultyWeight": 10
  },
  "Cocktail": {
    "count": 5,
    "lastSeenTime": 1586915139518
  },
  "sperren": {
    "count": 3,
    "lastSeenTime": 1586839270292
  },
  "Mount Everest": {
    "count": 4,
    "lastSeenTime": 1586957925802
  },
  "Puppenhaus": {
    "count": 3,
    "lastSeenTime": 1586879335186
  },
  "Jetski": {
    "count": 3,
    "lastSeenTime": 1586904223932
  },
  "Holzscheit": {
    "count": 6,
    "lastSeenTime": 1586890178799
  },
  "Zahnspange": {
    "count": 3,
    "lastSeenTime": 1586915974465,
    "difficulty": 1,
    "difficultyWeight": 5
  },
  "Florida": {
    "count": 3,
    "lastSeenTime": 1586896941227,
    "difficulty": 0.8,
    "difficultyWeight": 5
  },
  "Lorbeeren": {
    "count": 3,
    "lastSeenTime": 1586890464755
  },
  "Zirkusdirektor": {
    "count": 3,
    "lastSeenTime": 1586817544108,
    "difficulty": 0.5555555555555556,
    "difficultyWeight": 9
  },
  "Steuergerät": {
    "count": 2,
    "lastSeenTime": 1586784694118
  },
  "Shrek": {
    "count": 6,
    "lastSeenTime": 1586916342434
  },
  "copy": {
    "count": 1,
    "lastSeenTime": 1586783576683,
    "difficulty": 0.5,
    "difficultyWeight": 8
  },
  "blood": {
    "count": 1,
    "lastSeenTime": 1586783594582,
    "difficulty": 0.6,
    "difficultyWeight": 5
  },
  "Farn": {
    "count": 5,
    "lastSeenTime": 1586919722216
  },
  "eclipse": {
    "count": 1,
    "lastSeenTime": 1586783732447,
    "difficulty": 0.9411764705882353,
    "difficultyWeight": 34
  },
  "palm tree": {
    "count": 1,
    "lastSeenTime": 1586783812435,
    "difficulty": 0.5,
    "difficultyWeight": 12
  },
  "AC/DC": {
    "count": 5,
    "lastSeenTime": 1586958154127
  },
  "Schlafanzug": {
    "count": 2,
    "lastSeenTime": 1586813334338
  },
  "Teig": {
    "count": 8,
    "lastSeenTime": 1586919287618
  },
  "Kegel": {
    "count": 3,
    "lastSeenTime": 1586954910317
  },
  "Hamburger": {
    "count": 4,
    "lastSeenTime": 1586914590565,
    "difficulty": 0.8928571428571429,
    "difficultyWeight": 28
  },
  "Stecknadel": {
    "count": 6,
    "lastSeenTime": 1586907609327
  },
  "Invasion": {
    "count": 3,
    "lastSeenTime": 1586873772670
  },
  "klebrig": {
    "count": 3,
    "lastSeenTime": 1586895452509
  },
  "Flamingo": {
    "count": 4,
    "lastSeenTime": 1586908104595,
    "difficulty": 0.5714285714285714,
    "difficultyWeight": 7
  },
  "Orange": {
    "count": 4,
    "lastSeenTime": 1586890743822,
    "difficulty": 0.9,
    "difficultyWeight": 40
  },
  "Ball": {
    "count": 2,
    "lastSeenTime": 1586840879616
  },
  "Becken": {
    "count": 5,
    "lastSeenTime": 1586958547642,
    "difficulty": 1,
    "difficultyWeight": 8
  },
  "Verpackung": {
    "count": 3,
    "lastSeenTime": 1586815146910
  },
  "Bogenschütze": {
    "count": 3,
    "lastSeenTime": 1586920143399
  },
  "versöhnen": {
    "count": 6,
    "lastSeenTime": 1586959830012
  },
  "Rampe": {
    "count": 3,
    "lastSeenTime": 1586903111992
  },
  "Bronze": {
    "count": 5,
    "lastSeenTime": 1586959585641
  },
  "Wohlstand": {
    "count": 4,
    "lastSeenTime": 1586914991061
  },
  "Dessert": {
    "count": 3,
    "lastSeenTime": 1586903011803
  },
  "Bowling": {
    "count": 4,
    "lastSeenTime": 1586914615571
  },
  "Animation": {
    "count": 4,
    "lastSeenTime": 1586840045031
  },
  "Fitness Trainer": {
    "count": 6,
    "lastSeenTime": 1586956462502
  },
  "Hydrant": {
    "count": 4,
    "lastSeenTime": 1586959342619,
    "difficulty": 1,
    "difficultyWeight": 1
  },
  "Alligator": {
    "count": 6,
    "lastSeenTime": 1586919958324
  },
  "Atom": {
    "count": 5,
    "lastSeenTime": 1586959076238
  },
  "farbenblind": {
    "count": 2,
    "lastSeenTime": 1586814501266
  },
  "erfrieren": {
    "count": 1,
    "lastSeenTime": 1586805684110
  },
  "Tails": {
    "count": 1,
    "lastSeenTime": 1586805694893
  },
  "Staudamm": {
    "count": 2,
    "lastSeenTime": 1586901547884
  },
  "Pfütze": {
    "count": 3,
    "lastSeenTime": 1586895782227
  },
  "Engel": {
    "count": 2,
    "lastSeenTime": 1586956582710
  },
  "Wunderlampe": {
    "count": 4,
    "lastSeenTime": 1586869345288,
    "difficulty": 1,
    "difficultyWeight": 7
  },
  "Fechten": {
    "count": 4,
    "lastSeenTime": 1586954989864
  },
  "Gas": {
    "count": 5,
    "lastSeenTime": 1586920226877
  },
  "Kunde": {
    "count": 4,
    "lastSeenTime": 1586955456078
  },
  "Zelda": {
    "count": 2,
    "lastSeenTime": 1586880631457
  },
  "Atomuhr": {
    "count": 3,
    "lastSeenTime": 1586873715364
  },
  "Realität": {
    "count": 4,
    "lastSeenTime": 1586839250104
  },
  "Feuerwache": {
    "count": 5,
    "lastSeenTime": 1586959304612
  },
  "Schnee": {
    "count": 5,
    "lastSeenTime": 1586958769937,
    "difficulty": 0.7777777777777778,
    "difficultyWeight": 9
  },
  "Altweibersommer": {
    "count": 3,
    "lastSeenTime": 1586859631642
  },
  "Asterix": {
    "count": 4,
    "lastSeenTime": 1586959557107
  },
  "Bingo": {
    "count": 7,
    "lastSeenTime": 1586954925078
  },
  "Rosine": {
    "count": 5,
    "lastSeenTime": 1586919750850
  },
  "Koala": {
    "count": 5,
    "lastSeenTime": 1586958350172
  },
  "Propeller": {
    "count": 6,
    "lastSeenTime": 1586915200929
  },
  "hüpfen": {
    "count": 6,
    "lastSeenTime": 1586959753084
  },
  "Blatt": {
    "count": 2,
    "lastSeenTime": 1586920407752
  },
  "Asche": {
    "count": 2,
    "lastSeenTime": 1586958413299
  },
  "Tau": {
    "count": 4,
    "lastSeenTime": 1586908368588
  },
  "Seewolf": {
    "count": 2,
    "lastSeenTime": 1586955999033
  },
  "Minigolf": {
    "count": 2,
    "lastSeenTime": 1586815692372
  },
  "smell": {
    "count": 1,
    "lastSeenTime": 1586813203038,
    "difficulty": 0.45454545454545453,
    "difficultyWeight": 11
  },
  "sunglasses": {
    "count": 1,
    "lastSeenTime": 1586813258085,
    "difficulty": 0.7272727272727273,
    "difficultyWeight": 11
  },
  "waist": {
    "count": 1,
    "lastSeenTime": 1586813304541,
    "difficulty": 0.2222222222222222,
    "difficultyWeight": 9
  },
  "compass": {
    "count": 1,
    "lastSeenTime": 1586813320650,
    "difficulty": 1,
    "difficultyWeight": 9
  },
  "drinnen": {
    "count": 3,
    "lastSeenTime": 1586874931565
  },
  "Sprecher": {
    "count": 3,
    "lastSeenTime": 1586890093454
  },
  "flashlight": {
    "count": 1,
    "lastSeenTime": 1586813394246,
    "difficulty": 0.4444444444444444,
    "difficultyWeight": 9
  },
  "Fußende": {
    "count": 5,
    "lastSeenTime": 1586915406529
  },
  "Rasierer": {
    "count": 3,
    "lastSeenTime": 1586901284755,
    "difficulty": 0.9375,
    "difficultyWeight": 32
  },
  "Turm": {
    "count": 1,
    "lastSeenTime": 1586813676447
  },
  "Bürgermeister": {
    "count": 2,
    "lastSeenTime": 1586816602699,
    "difficulty": 1,
    "difficultyWeight": 5
  },
  "Esel": {
    "count": 4,
    "lastSeenTime": 1586958350172,
    "difficulty": 1,
    "difficultyWeight": 8
  },
  "Meme": {
    "count": 2,
    "lastSeenTime": 1586833304748,
    "difficulty": 0.625,
    "difficultyWeight": 16
  },
  "Horizont": {
    "count": 2,
    "lastSeenTime": 1586914638481
  },
  "schnorcheln": {
    "count": 4,
    "lastSeenTime": 1586903429291
  },
  "Abschluss": {
    "count": 2,
    "lastSeenTime": 1586868207231
  },
  "Belgien": {
    "count": 4,
    "lastSeenTime": 1586909070561,
    "difficulty": 0.5555555555555556,
    "difficultyWeight": 9
  },
  "Grille": {
    "count": 1,
    "lastSeenTime": 1586814301944
  },
  "Meeresfrüchte": {
    "count": 7,
    "lastSeenTime": 1586904308260
  },
  "sea": {
    "count": 1,
    "lastSeenTime": 1586814732814,
    "difficulty": 1,
    "difficultyWeight": 2
  },
  "Jesus Christ": {
    "count": 1,
    "lastSeenTime": 1586814734854
  },
  "cell": {
    "count": 1,
    "lastSeenTime": 1586814734854
  },
  "Schach": {
    "count": 5,
    "lastSeenTime": 1586959443697,
    "difficulty": 0.8484848484848485,
    "difficultyWeight": 33
  },
  "Schädel": {
    "count": 3,
    "lastSeenTime": 1586880722993,
    "difficulty": 0.7333333333333333,
    "difficultyWeight": 15
  },
  "Delfin": {
    "count": 1,
    "lastSeenTime": 1586815327708,
    "difficulty": 0.7142857142857143,
    "difficultyWeight": 7
  },
  "mother": {
    "count": 1,
    "lastSeenTime": 1586815477720,
    "difficulty": 0.8444444444444444,
    "difficultyWeight": 45
  },
  "wart": {
    "count": 1,
    "lastSeenTime": 1586815533308,
    "difficulty": 1,
    "difficultyWeight": 10
  },
  "cone": {
    "count": 1,
    "lastSeenTime": 1586815591945,
    "difficulty": 0.9,
    "difficultyWeight": 10
  },
  "sneeze": {
    "count": 1,
    "lastSeenTime": 1586815651308,
    "difficulty": 0.6923076923076923,
    "difficultyWeight": 13
  },
  "rice": {
    "count": 1,
    "lastSeenTime": 1586815683909,
    "difficulty": 0.6923076923076923,
    "difficultyWeight": 13
  },
  "Anfänger": {
    "count": 3,
    "lastSeenTime": 1586955999033
  },
  "kochen": {
    "count": 2,
    "lastSeenTime": 1586955569169,
    "difficulty": 1,
    "difficultyWeight": 3
  },
  "westlich": {
    "count": 1,
    "lastSeenTime": 1586816279864
  },
  "Burgruine": {
    "count": 1,
    "lastSeenTime": 1586816334630
  },
  "Gandhi": {
    "count": 3,
    "lastSeenTime": 1586955334605
  },
  "Haarschnitt": {
    "count": 5,
    "lastSeenTime": 1586914820916
  },
  "Planet": {
    "count": 3,
    "lastSeenTime": 1586909201701,
    "difficulty": 0.8,
    "difficultyWeight": 20
  },
  "Nessie": {
    "count": 2,
    "lastSeenTime": 1586839320824
  },
  "Kim Kardashian": {
    "count": 3,
    "lastSeenTime": 1586859216557
  },
  "Widerstand": {
    "count": 2,
    "lastSeenTime": 1586839211738
  },
  "Lebkuchen": {
    "count": 3,
    "lastSeenTime": 1586897469693
  },
  "Genie": {
    "count": 3,
    "lastSeenTime": 1586900436732
  },
  "Tyrannosaurus Rex": {
    "count": 2,
    "lastSeenTime": 1586956352310
  },
  "grün": {
    "count": 4,
    "lastSeenTime": 1586907557164
  },
  "Flash": {
    "count": 2,
    "lastSeenTime": 1586909155516
  },
  "Zorro": {
    "count": 3,
    "lastSeenTime": 1586859604461
  },
  "Kommunismus": {
    "count": 3,
    "lastSeenTime": 1586915919352
  },
  "Camping": {
    "count": 3,
    "lastSeenTime": 1586900977668,
    "difficulty": 1,
    "difficultyWeight": 9
  },
  "Vogelhaus": {
    "count": 2,
    "lastSeenTime": 1586890938315,
    "difficulty": 0.5714285714285714,
    "difficultyWeight": 7
  },
  "Ritter": {
    "count": 3,
    "lastSeenTime": 1586915741689
  },
  "Brocken": {
    "count": 5,
    "lastSeenTime": 1586958253918
  },
  "Krankheit": {
    "count": 3,
    "lastSeenTime": 1586914119532,
    "difficulty": 1,
    "difficultyWeight": 8
  },
  "Teelöffel": {
    "count": 5,
    "lastSeenTime": 1586921104275,
    "difficulty": 1,
    "difficultyWeight": 3
  },
  "Kaltwachsstreifen": {
    "count": 4,
    "lastSeenTime": 1586916127462
  },
  "Keim": {
    "count": 2,
    "lastSeenTime": 1586909311071
  },
  "Jahrbuch": {
    "count": 1,
    "lastSeenTime": 1586827501149,
    "difficulty": 1,
    "difficultyWeight": 1
  },
  "verrückt": {
    "count": 1,
    "lastSeenTime": 1586827603168
  },
  "Gewalt": {
    "count": 4,
    "lastSeenTime": 1586916540283
  },
  "Hüpfkästchen": {
    "count": 3,
    "lastSeenTime": 1586870007658
  },
  "Kerzenständer": {
    "count": 3,
    "lastSeenTime": 1586860109145
  },
  "Zauberstab": {
    "count": 4,
    "lastSeenTime": 1586956324955
  },
  "Telefonkabel": {
    "count": 3,
    "lastSeenTime": 1586915048330
  },
  "Filmemacher": {
    "count": 3,
    "lastSeenTime": 1586958611285,
    "difficulty": 1,
    "difficultyWeight": 2
  },
  "Spielplatz": {
    "count": 2,
    "lastSeenTime": 1586904412233
  },
  "Ohrenschmalz": {
    "count": 1,
    "lastSeenTime": 1586839211738,
    "difficulty": 0.8888888888888888,
    "difficultyWeight": 9
  },
  "Kerzenleuchter": {
    "count": 2,
    "lastSeenTime": 1586958572090
  },
  "Winter": {
    "count": 2,
    "lastSeenTime": 1586915225560,
    "difficulty": 0.6153846153846154,
    "difficultyWeight": 13
  },
  "Chef": {
    "count": 5,
    "lastSeenTime": 1586920932924
  },
  "iPhone": {
    "count": 5,
    "lastSeenTime": 1586958596468
  },
  "Chrome": {
    "count": 1,
    "lastSeenTime": 1586840244704,
    "difficulty": 0.75,
    "difficultyWeight": 8
  },
  "Speichel": {
    "count": 2,
    "lastSeenTime": 1586873812676
  },
  "Ratte": {
    "count": 2,
    "lastSeenTime": 1586889601390
  },
  "Grill": {
    "count": 2,
    "lastSeenTime": 1586896066306
  },
  "Prisma": {
    "count": 2,
    "lastSeenTime": 1586920346900
  },
  "Orca": {
    "count": 1,
    "lastSeenTime": 1586859539297
  },
  "Sicherung": {
    "count": 5,
    "lastSeenTime": 1586920397655
  },
  "Sandbank": {
    "count": 2,
    "lastSeenTime": 1586869853196
  },
  "Nachthemd": {
    "count": 1,
    "lastSeenTime": 1586868911645
  },
  "Bauch": {
    "count": 1,
    "lastSeenTime": 1586869237926,
    "difficulty": 0.8421052631578947,
    "difficultyWeight": 19
  },
  "Johnny Bravo": {
    "count": 2,
    "lastSeenTime": 1586874716360
  },
  "Stachelrochen": {
    "count": 3,
    "lastSeenTime": 1586956215415
  },
  "Dreck": {
    "count": 1,
    "lastSeenTime": 1586874073864
  },
  "Taser": {
    "count": 2,
    "lastSeenTime": 1586907375887,
    "difficulty": 1,
    "difficultyWeight": 9
  },
  "Metall": {
    "count": 4,
    "lastSeenTime": 1586955010181
  },
  "Doktor": {
    "count": 3,
    "lastSeenTime": 1586955763543
  },
  "Schriftsteller": {
    "count": 6,
    "lastSeenTime": 1586958934411
  },
  "laufen": {
    "count": 1,
    "lastSeenTime": 1586880438207
  },
  "Alarm": {
    "count": 2,
    "lastSeenTime": 1586959269543
  },
  "Bankier": {
    "count": 2,
    "lastSeenTime": 1586916756124
  },
  "Stinktier": {
    "count": 3,
    "lastSeenTime": 1586920275708
  },
  "Tresorraum": {
    "count": 1,
    "lastSeenTime": 1586890625702
  },
  "Nichtschwimmer": {
    "count": 1,
    "lastSeenTime": 1586890743822
  },
  "Opernhaus Sydney": {
    "count": 2,
    "lastSeenTime": 1586907532918
  },
  "jagen": {
    "count": 2,
    "lastSeenTime": 1586902702172
  },
  "Autor": {
    "count": 2,
    "lastSeenTime": 1586959840559
  },
  "Jackie Chan": {
    "count": 1,
    "lastSeenTime": 1586900412247
  },
  "Reflexion": {
    "count": 2,
    "lastSeenTime": 1586919650407
  },
  "Zweig": {
    "count": 2,
    "lastSeenTime": 1586955899639
  },
  "Experiment": {
    "count": 1,
    "lastSeenTime": 1586908921510
  },
  "Badezimmer": {
    "count": 3,
    "lastSeenTime": 1586959738785
  },
  "Spüllappen": {
    "count": 1,
    "lastSeenTime": 1586915543878
  },
  "Magnet": {
    "count": 1,
    "lastSeenTime": 1586915703672
  },
  "Fühler": {
    "count": 2,
    "lastSeenTime": 1586955098414
  },
  "Eisbär": {
    "count": 1,
    "lastSeenTime": 1586920361074
  },
  "Grabmal": {
    "count": 1,
    "lastSeenTime": 1586958636395
  }
}
`
