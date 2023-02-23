import okhttp3.ResponseBody
import retrofit2.Response
import retrofit2.http.Header
import retrofit2.http.POST
import retrofit2.http.Query

interface ApiInterface {

    @POST("/nic/update")
    suspend fun post(@Query("hostname") hostname: String,
                     @Header("Authorization") auth: String): Response<ResponseBody>
}