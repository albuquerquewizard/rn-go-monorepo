import { Text, View } from "react-native";
import "./global.css"

export default function Index() {
  return (
    <View className="flex-1 justify-center items-center bg-blue-100">
      <Text className="text-2xl font-bold text-blue-800 mb-4">
        NativeWind is working lessgo!!!
      </Text>
      <Text className="text-lg text-gray-600 text-center px-4">
        If you can see this styled text with blue background, NativeWind is working! 🎉
      </Text>
    </View>
  );
}
