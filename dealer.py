with open("assetIdInfo") as fp:
    for line in fp:
        line = line.split(",")
        # print(line[0].split("("))
        # print(line[1])
        # print(line[2])
        # print(line[3])
        # print(line[4])
        # print(line[5])
        # print(line[6].replace(")", ""))

        id = line[0].split("(")[0].replace(": ", "")
        name = line[0].split("(")[1]
        chainlinkFeed = line[1]
        minLeverage = line[2]
        maxLeverage = line[3]
        feeMultiplier = line[4]
        baseFundingRate = line[5]
        pythnetId = line[6].replace(")", "").replace('\n', '')

        print(f"TigrisIDToAssetInfo[{id}] = AssetInfo " + "{")
        print(f"Name: \"{name}\",")
        print(f"ChainlinkFeed: \"{chainlinkFeed}\",")
        print(f"MinLeverage: \"{minLeverage}\",")
        print(f"MaxLeverage: \"{maxLeverage}\",")
        print(f"FeeMultiplier: \"{feeMultiplier}\",")
        print(f"BaseFundingRate: \"{baseFundingRate}\",")
        print(f"PythnetId: \"{pythnetId}\",")
        print("}")
        print("")

        # print(id, name, chainlinkFeed, minLeverage, maxLeverage,
        #       feeMultiplier, baseFundingRate, pythnetId)

        # name            string
        # chainlinkFeed   string
        # minLeverage     string
        # maxLeverage     string
        # feeMultiplier   string
        # baseFundingRate string
        # pythnetId       string
