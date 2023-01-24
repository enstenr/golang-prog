package main

import (
	"fmt"
	"time"
	"github.com/enstenr/common/dao"
	"github.com/enstenr/common/connection"
	"github.com/enstenr/customtypes"
	"github.com/spf13/viper"
)
func main (){
	 
	 
	 LoadProperties("stage");
	  
	  LoadProperties("canary")
	  smap:=make(map[string]customtypes.MetadataTree)
	  cmap:=make(map[string]customtypes.MetadataTree)
	  itemArrayStage:=make([]customtypes.MetadataTree,0)
	  itemArrayStage=dao.FetchTree("stage")
	  for _,sObj :=range itemArrayStage{
		//fmt.Println(value.MetadataTreeConfigurationId,value.Hash, value.Modified_date)
		smap[sObj.MetadataTreeConfigurationId]=sObj

	  }
	 
	  itemArrayCanary:=make([]customtypes.MetadataTree,0)
	  itemArrayCanary=dao.FetchTree("canary")
	  for _,cObj :=range itemArrayCanary{
		//fmt.Println(value.MetadataTreeConfigurationId,value.Hash, value.Modified_date)
		cmap[cObj.MetadataTreeConfigurationId]=cObj

	  }
	  treesInStageNotinCanary(smap,cmap)

	  var index=1
	  outputArray:=make([]customtypes.MetadataTreeDifference,0)
	  fmt.Println(" Metadata Configuration Id\tCanary Hash\tStage Hash\tCanary  Modified Date\tStage  Modified Date\t Canary RePublishStatus\t Stage RePublishStatus ")
	  for _,cObj :=range itemArrayCanary{
		sObj:=smap[cObj.MetadataTreeConfigurationId]
		format:="2012-01-26T22:58:28.453207Z"
		timeObj1,_:=time.Parse(format,sObj.Modified_date)
		timeObj2,_:=time.Parse(format,cObj.Modified_date)
		if(timeObj1.Before(timeObj2)){
			//fmt.Print(timeObj1.Before(timeObj2))
	  }
		//fmt.Println(cObj.MetadataTreeConfigurationId,sObj.MetadataTreeConfigurationId, cObj.Hash,cObj.Hash)
		if( cObj.Hash == sObj.Hash){
			outputObj:=customtypes.MetadataTreeDifference{}
			outputObj.MetadataTreeConfigurationId1=cObj.MetadataTreeConfigurationId
			outputObj.MetadataTreeConfigurationId2=sObj.MetadataTreeConfigurationId
			outputObj.Hash1=cObj.Hash 
			outputObj.Hash2=sObj.Hash
			outputObj.Modified_date1=cObj.Modified_date
			outputObj.Modified_date2=sObj.Modified_date
			outputObj.RePublishStatus1=cObj.RePublishStatus
			outputObj.RePublishStatus2=sObj.RePublishStatus

			outputArray=append(outputArray,outputObj )
		//fmt.Println(index, cObj.MetadataTreeConfigurationId,cObj.Hash,sObj.Hash, cObj.Modified_date,sObj.Modified_date,cObj.RePublishStatus,sObj.RePublishStatus)
		index++
	  }
	}
	 connection.WriteTreeToCSV(outputArray,"same_hash_comparison.csv")

}

func treesInStageNotinCanary(smap map[string]customtypes.MetadataTree, cmap map[string]customtypes.MetadataTree){
	fmt.Println(" Tree not in canary but available in Stage")
	for skey,_ := range smap{
		_,cok:=cmap[skey]
		if(!cok){
			
			fmt.Println(skey)
		}
	}
}
var globalViperObj *viper.Viper;
func LoadProperties(env string)( *viper.Viper){
	viperObj1 :=viper.New();
	viperObj1.SetConfigName(env)
	viperObj1.SetConfigType("env")
	viperObj1.AddConfigPath(".")
	viperObj1.ReadInConfig()
	globalViperObj=viperObj1
	return viperObj1
}